package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	originalPhoneNumbers = []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
)

func formatNumber(number string) string {

	re, err := regexp.Compile(`\d`)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(re.FindAllString(number, -1), "")
}

func main() {
	fmt.Println(formatNumber("(123)456-7892"))
	// open the db
	db, err := sql.Open("sqlite3", "phone.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// create the table
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS phone_numbers (id INTEGER PRIMARY KEY AUTOINCREMENT, phone_number VARCHAR(64)); "); err != nil {
		log.Fatalf("Error while executing query %v", err)
	}
	rowsCount, _ := db.Query("Select count(1) from phone_numbers;")
	var count int
	for rowsCount.Next() {
		rowsCount.Scan(&count)
	}
	rowsCount.Close()
	if count == 0 {
		// prepare the insert query
		stmt, err := db.Prepare("INSERT into phone_numbers (phone_number) Values (?);")
		if err != nil {
			log.Fatalf("Failed to perparing insert query: %v", err)
		}
		// insert the values from the array
		for _, phoneNumber := range originalPhoneNumbers {
			if _, err := stmt.Exec(phoneNumber); err != nil {
				log.Fatalf("Error while executing insert command %v", err)
			}
		}
	}

	// listing the phone numbers from the db
	rows, err := db.Query("Select id, phone_number from phone_numbers;")
	if err != nil {
		log.Fatalf("Failed to executing select query: %v", err)
	}
	defer rows.Close()

	type rowType struct {
		id           int16
		phone_number string
	}
	var rowsToBeFormatted []rowType
	for rows.Next() {
		var (
			id           sql.NullInt16
			phone_number string
		)
		if err := rows.Scan(&id, &phone_number); err != nil {
			log.Fatal(err)
		}
		formattedNumber := formatNumber(phone_number)

		if formattedNumber != phone_number && id.Valid {
			rowDetail := rowType{id: id.Int16, phone_number: formattedNumber}
			rowsToBeFormatted = append(rowsToBeFormatted, rowDetail)
			fmt.Println("updated")
		}
	}
	// preparing update statement and formatting numbers
	updtStmt, err := db.Prepare("UPDATE phone_numbers SET phone_number = ? WHERE id = ?;")
	if err != nil {
		log.Fatalf("Failed to executing update query: %v", err)
	}

	for _, row := range rowsToBeFormatted {
		if _, err := updtStmt.Exec(row.phone_number, row.id); err != nil {
			log.Fatal(err)
		}
	}

	// selecting distinct
	distinctRows, err := db.Query("Select distinct(phone_number) from phone_numbers;")
	if err != nil {
		log.Fatalf("Failed to executing select query: %v", err)
	}
	defer distinctRows.Close()
	for distinctRows.Next() {
		var (
			phone_number string
		)
		if err := distinctRows.Scan(&phone_number); err != nil {
			log.Fatal(err)
		}
		fmt.Println(phone_number)
	}

}
