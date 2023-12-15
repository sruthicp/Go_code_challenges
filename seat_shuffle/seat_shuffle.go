/*
Table: Seat
+-------------+---------+
| Column Name | Type |
+-------------+---------+
| id | int |
| student | varchar |
+-------------+---------+
id is the primary key (unique value) column for this table.
Each row of this table indicates the name and the ID of a student.
id is a continuous increment.

Write a solution to swap the seat id of every two consecutive students. If the number of students is odd,
the id of the last student is not swapped.
Return the result table ordered by id in ascending order.
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Seat represents the Seat table
type Seat struct {
	ID      int
	Student string
}

func main() {
	// Database connection parameters
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "postgres"

	// Create the database connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the Seat table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS Seat (
		id SERIAL PRIMARY KEY,
		student VARCHAR
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insert data into the Seat table
	insertDataSQL := `
	INSERT INTO Seat (student) VALUES
	('Abbot'),
	('Doris'),
	('Emerson'),
	('Green'),
	('Jeames');
	`

	_, err = db.Exec(insertDataSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Update the id of student, if id is odd and not the last one then id = id +1
	// if id is even id = id-1. Then return the values in ascending order.
	swapSeatsSQL := `
	SELECT
		CASE
			WHEN id % 2 = 1 AND id < (SELECT MAX(id) FROM Seat) THEN id + 1
			WHEN id % 2 = 0 THEN id - 1
			ELSE id
		END 
	AS id,student
	FROM Seat
	ORDER BY id;
	`

	rows, err := db.Query(swapSeatsSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the result
	fmt.Println("Output:")
	fmt.Printf("+----+---------+\n")
	fmt.Printf("| id | student |\n")
	fmt.Printf("+----+---------+\n")

	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.ID, &seat.Student)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(" %d   |  %s  \n", seat.ID, seat.Student)
	}

	fmt.Printf("+----+---------+\n")

	// Drop the Seat table
	dropTableSQL := "DROP TABLE Seat;"
	_, err = db.Exec(dropTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}
