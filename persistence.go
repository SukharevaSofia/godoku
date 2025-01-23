package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// makes sure that all tables and the db itself exist
func db_manager() {
	fmt.Println("opening the db...")
	db, err := sql.Open("sqlite3", "godoku.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	log.Println(version)
	// Create tables for fields and rows
	const create_fields_query = FIELDS_TABLE
	_, err = db.Exec(create_fields_query)
	if err != nil {
		log.Println("Fields database creation failed:")
		log.Println(err)
		panic(err)
	}

	// Create
	const create_rows_query = ROWS_TABLE
	_, err = db.Exec(create_rows_query)
	if err != nil {
		log.Println("Rows database creation failed:")
		log.Println(err)
		panic(err)
	}
}

// adds field to the db and returns the id of a newly added field
func save_field(full_field, play_field, hints_loc [81]int, 
  save_name string, is_solved bool) (int64, int64, int64) {
	// opening the db
	db, err := sql.Open("sqlite3", "godoku.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()

	fmt.Println("updating the db...")

	// adding the info into fields table
  full_field_id := save_full_field(db, save_name, is_solved, full_field)
  play_field_id := save_play_field(db, save_name, is_solved, play_field)
  hints_loc_id := save_hints_loc(db, save_name, is_solved, hints_loc)

  return full_field_id, play_field_id, hints_loc_id
}

func save_full_field(db *sql.DB, save_name string, is_solved bool, full_field [81]int) int64 {
	insert_field, err := db.Exec(INSERT_FIELDS, save_name, is_solved, true, false)

	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Printf("Inserted the %s field into database.\n", save_name)
	}

	field_id, _ := insert_field.LastInsertId()
	// adding the info into rows tables
	save_rows_of_field(db, field_id, full_field)

	return field_id
}

func save_play_field(db *sql.DB, save_name string, is_solved bool, play_field [81]int) int64 {
	insert_field, err := db.Exec(INSERT_FIELDS, save_name, is_solved, false, true)

	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Printf("Inserted the %s field into database.\n", save_name)
	}

	field_id, _ := insert_field.LastInsertId()
	// adding the info into rows table
	save_rows_of_field(db, field_id, play_field)

	return field_id
}

func save_hints_loc(db *sql.DB, save_name string, is_solved bool, hints_field [81]int) int64 {
	insert_field, err := db.Exec(INSERT_FIELDS, save_name, is_solved, false, false)

	if err != nil {
		log.Println("Error while inserting.")
		log.Println(err)
		panic(err)
	} else {
		log.Printf("Inserted the %s field into database.\n", save_name)
	}

	field_id, _ := insert_field.LastInsertId()
	save_rows_of_field(db, field_id, hints_field)
	// adding the info into rows table

	return field_id
}

func save_rows_of_field(db *sql.DB, field_id int64, values [81]int) {
	for row_id := 0; row_id < 9; row_id++ {
		_, err := db.Exec("INSERT INTO rows VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			field_id, row_id,
			values[row_id*9],
			values[row_id*9+1],
			values[row_id*9+2],
			values[row_id*9+3],
			values[row_id*9+4],
			values[row_id*9+5],
			values[row_id*9+6],
			values[row_id*9+7],
			values[row_id*9+8])
		if err != nil {
			log.Println("Error while inserting row.")
			log.Println(err)
			panic(err)
		} else{
      log.Printf("inserted rows to %d table", field_id)
    }
	}
}

func show_fields(show_obscured bool) {
	db, err := sql.Open("sqlite3", "godoku.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()
  
  var is_full, is_play bool
  if show_obscured{
    is_full = false
    is_play = true
  }else{
    is_full = true
    is_play = false
  }

	fields, err := db.Query("SELECT id, save_name, is_solved FROM fields " +
    "WHERE is_full_field = ? AND is_play_field = ?", is_full, is_play)
	if err != nil {
		log.Println("couldn't exec selecting from fields table")
		panic(err)
	}

	temp_field := field{}

	for fields.Next() {
		fields.Scan(&temp_field.Id, &temp_field.Save_name, &temp_field.Is_solved)
    fmt.Println("-----------------------------")
		fmt.Printf("field id: %d, name: %s, solved? %v\n",
			temp_field.Id, temp_field.Save_name, temp_field.Is_solved)

		rows, err := db.Query("SELECT * FROM rows WHERE field_id = ?", temp_field.Id)
		if err != nil {
			log.Println("couldn't Prepare statment for quering rows")
			panic(err)
		}

		var temp string
		row_num := 0
		for rows.Next() {
			rows.Scan(&temp, &temp,
				&temp_field.full_field_values[0+9*row_num],
				&temp_field.full_field_values[1+9*row_num],
				&temp_field.full_field_values[2+9*row_num],
				&temp_field.full_field_values[3+9*row_num],
				&temp_field.full_field_values[4+9*row_num],
				&temp_field.full_field_values[5+9*row_num],
				&temp_field.full_field_values[6+9*row_num],
				&temp_field.full_field_values[7+9*row_num],
				&temp_field.full_field_values[8+9*row_num])
			row_num++
		}
		print_field(temp_field.full_field_values)
	}
}
