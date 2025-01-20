package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// makes sure that all tables and the db itself exist
func db_manager() {
	fmt.Println("opening the db...")
	db, err := sql.Open("sqlite3", "godoku.db")

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
	// Create tables for fields and rows
	// TODO: add more constraints
	const create_fields_query = `CREATE TABLE IF NOT EXISTS fields
		(id INTEGER PRIMARY KEY UNIQUE NOT NULL, 
		save_name TEXT NOT  NULL, 
		is_solved BOOLEAN)`
	_, err = db.Exec(create_fields_query)
	if err != nil {
		log.Println("Fields database creation failed:")
		log.Println(err)
		panic(err)
	}

	// Create
	const create_rows_query = `CREATE TABLE IF NOT EXISTS rows
		(field_id INTEGER FOREIGT KEY NOT NULL, 
		row_id INTEGER NOT NULL, 
		val_1 INTEGER,
		val_2 INTEGER,
		val_3 INTEGER,
		val_4 INTEGER,
		val_5 INTEGER,
		val_6 INTEGER,
		val_7 INTEGER,
		val_8 INTEGER,
		val_9 INTEGER)`

	_, err = db.Exec(create_rows_query)
	if err != nil {
		log.Println("Rows database creation failed:")
		log.Println(err)
		panic(err)
	}
}

// adds field to the db and returns the id of a newly added field
func add_field(field []int, save_name string, is_solved bool) int64 {
	// opening the db
	db, err := sql.Open("sqlite3", "godoku.db")
	fmt.Println("adding to the db...")

	// adding the info into fields table
  insert_field, err := db.Exec("INSERT INTO fields (save_name, is_solved) VALUES (?, ?)", save_name, is_solved)
	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Printf("Inserted the %s field into database.\n", save_name)
	}
  
  field_id, _ := insert_field.LastInsertId()

	// adding the info into rows table
	for row_id := 0; row_id < 9; row_id++ {
		_, err = db.Exec("INSERT INTO rows VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			field_id, row_id,
			field[row_id*9],
			field[row_id*9+1],
			field[row_id*9+2],
			field[row_id*9+3],
			field[row_id*9+4],
			field[row_id*9+5],
			field[row_id*9+6],
			field[row_id*9+7],
			field[row_id*9+8])
		if err != nil {
			log.Println("Error while inserting row.")
			log.Println(err)
			panic(err)
		} else {
			log.Printf("Inserted the %s field into database.\n", save_name)
		}
	}
	fmt.Println("successfuly added!")

	return field_id
}

func show_fields() {
	db, err := sql.Open("sqlite3", "godoku.db")
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fields, err := db.Query("SELECT id, save_name, is_solved FROM fields")
	if err != nil {
		log.Println("couldn't exec selecting from fields table")
		panic(err)
	}

	temp_field := field{}

	for fields.Next() {
		fields.Scan(&temp_field.Id, &temp_field.Save_name, &temp_field.Is_solved)
		fmt.Printf("|-|-|-|-|-|\n field id: %d, name: %s, solved? %v\n",
			temp_field.Id, temp_field.Save_name, temp_field.Is_solved)

		rows, err := db.Query("SELECT * FROM rows WHERE field_id = ?", temp_field.Id)
		if err != nil {
			log.Println("couldn't Prepare statment for quering rows")
			panic(err)
		}

		var temp string
		var row_id int
		for rows.Next() {

			rows.Scan(&temp)
			row_id, _ = strconv.Atoi(temp)
			for i := 0; i < 9; i++ {
				rows.Scan(&temp_field.values[i+9*row_id])
			}
		}
		printField(temp_field.values)
	}
}
