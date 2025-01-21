package main

import (
	"fmt"
)

const (
	DIFFICULTY_CHOICE = "\nPlease, choose the level of difficulty:\n" +
		"1: Easy\n" +
		"2: Medium\n" +
		"3: Hard\n" +
		"> "

	GEN_OR_SHOW = "\n what should we do?\n" +
		"1: generate a field \n" +
		"2: show saved fields\n" +
		"> "

	SHOW_FULL_OR_HINT = "\n Show fields with obscured values?\n" +
		"1: yes please \n" +
		"2: no, show all the values\n" +
		"> "

	PLAY_CHOICE = "\n Wanna play on the generated field?\n" +
		"1: yes \n" +
		"2: no\n" +
		"> "

	SAVE_CHOICE = "\n Wanna save the field?\n" +
		"1: yes \n" +
		"2: no\n" +
		"> "

	FIELDS_TABLE = "CREATE TABLE IF NOT EXISTS fields" +
		"(id INTEGER PRIMARY KEY UNIQUE NOT NULL, " +
		"save_name TEXT NOT NULL, " +
		"is_solved BOOLEAN," +
		"is_full_field BOOLEAN," +
		"is_play_field BOOLEAN)"

	ROWS_TABLE = "CREATE TABLE IF NOT EXISTS rows" +
		"(field_id INTEGER NOT NULL, " +
		"row_id INTEGER NOT NULL, " +
		"val_1 INTEGER NOT NULL, " +
		"val_2 INTEGER NOT NULL, " +
		"val_3 INTEGER NOT NULL, " +
		"val_4 INTEGER NOT NULL, " +
		"val_5 INTEGER NOT NULL, " +
		"val_6 INTEGER NOT NULL, " +
		"val_7 INTEGER NOT NULL, " +
		"val_8 INTEGER NOT NULL, " +
		"val_9 INTEGER NOT NULL) "

	INSERT_FIELDS = "INSERT INTO fields (save_name, is_solved," +
		"is_full_field, is_play_field) VALUES (?, ?, ?, ?)"

	BAD_INPUT      = "incorrect input value"
	CLEAN_TERMINAL = "\033[H\033[2J"
)

// TODO: add a func to check if fields equal
type field struct {
	Id                int64
	Save_name         string
	Is_solved         bool
	full_field_values [81]int
	play_field_values [81]int
	hints_indexes     map[int]bool
}

func print_field(field [81]int) {
	for i, el := range field {

		if i%9 == 0 {
			fmt.Print("\n")
		}

		if (i%3 == 0) && ((i % 9) != 0) {
			fmt.Print("|")
		}
		if el == 0 {
			fmt.Print(" · ")
		} else {
			fmt.Printf(" %d ", el)
		}
		if (i == 26) || (i == 53) {
			fmt.Print("\n---------+---------+--------")
		}
	}
	fmt.Println("")
}
