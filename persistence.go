package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func db_manager() {
	db, err := sql.Open("sqlite3", "godoku.db")

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
	// Create table
  //TODO: add more constraints
	const create_query = `CREATE TABLE IF NOT EXISTS fields
		(id INTEGER PRIMARY KEY UNIQUE NOT NULL, 
		save_name TEXT NOT  NULL, 
		is_solved BOOLEAN)`
	_, err = db.Exec(create_query)
	if err != nil {
		log.Println("Database creation failed:")
		log.Println(err)
		panic(err)
	}
	// Create

	_, err = db.Exec("INSERT INTO fields (save_name, is_solved) VALUES (?, ?)", "Testing save", true)
	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Println("Inserted the field into database.")
	}

	// Read
	rows, err := db.Query("SELECT id, save_name, is_solved FROM fields")
	if err != nil {
		panic(err)
	}

	temp_field := field{}
	for rows.Next() {
		rows.Scan(&temp_field.Id, &temp_field.Save_name, &temp_field.Is_solved)
		log.Printf("ID:%d, Save name:%s, Is solved:%t\n", temp_field.Id,
			temp_field.Save_name, temp_field.Is_solved)
	}
}

func add_field(save_name string, is_solved bool) {
	db, err := sql.Open("sqlite3", "godoku.db")

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

  	_, err = db.Exec("INSERT INTO fields (save_name, is_solved) VALUES (?, ?)", save_name, is_solved)
	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Println("Inserted the field into database.")
	}



}
