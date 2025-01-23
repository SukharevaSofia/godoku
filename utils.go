package main

import (
	"fmt"
)

const (
	DIFFICULTY_CHOICE = 
`.-------------------------------------.
| Please, choose the level            |
| of difficulty:                      |
| 1: easy                             |
| 2: medium                           |
| 3: hard                             |
'-------------------------------------'
> `

	GEN_OR_SHOW = 
`.-------------------------------------.
| Welcome to godoku!                  |
| What would you want to do?          |
| 1: generate a field                 |
| 2: show saved fields                |
'-------------------------------------'
> `


	SHOW_FULL_OR_HINT = 
`.-------------------------------------.
| Show fields with obscured values?   |
| 1: yes please                       |
| 2: no, show all the values          |
'-------------------------------------'
> `

	PLAY_CHOICE = 
`.-------------------------------------.
| Wanna play the generated field?     |
| 1: yes                              |
| 2: no                               |
'-------------------------------------'
> `

	SAVE_CHOICE = 
`.-------------------------------------.
| Save the field?                     |
| 1: yes                              |
| 2: no                               |
'-------------------------------------'
> `

	WINNER = 
  `
 / \------------------------------------, 
 \_,|                                   | 
    |    Congratulations! You've won!   |
    |    Press any button to continue.  |
    |  ,----------------------------------
    \_/_________________________________/ `

	INSTRUCTIONS = 
`

.--------------------HOW-TO-PLAY--------------------.
| - move around using arrows ←↑→↓                   |
| - to change a value in a cell, press "enter" and  |
| type the desired value for that cell.             |
| - to delete a value, put 0 in a cell.             |
| - to finish and save session, press "c".          |
'---------------------------------------------------'
  `

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

func bad_input() {
	fmt.Println(BAD_INPUT)
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

func print_field_with_coursor(field [81]int, cursor int) {
	var is_choosen, is_lit_col, is_lit_row bool

	lit_column := cursor % 9
	lit_row := cursor / 9
	for i, el := range field {
		is_choosen = (i == cursor)
		is_lit_col = (i%9 == lit_column)
		is_lit_row = (i/9 == lit_row)

		if i%9 == 0 {
			fmt.Print("\n")
		}

		if (i%3 == 0) && ((i % 9) != 0) {
			fmt.Print("|")
		}
		if el == 0 {
			if is_choosen {
				fmt.Print("\033[45m . \033[m")
			} else if is_lit_row || is_lit_col {
				fmt.Print("\033[35m . \033[m")
			} else {
				fmt.Print("\033[0m . \033[m")
			}
		} else {
			if is_choosen {
				fmt.Printf("\033[45m %d \033[m", el)
			} else if is_lit_row || is_lit_col {
				fmt.Printf("\033[35m %d \033[m", el)
			} else {
				fmt.Printf("\033[0m %d \033[m", el)
			}
		}
		if (i == 26) || (i == 53) {
			fmt.Print("\n---------+---------+--------")
		}
	}
	fmt.Println("")
}

func equal_fields(input_field, correct_field [81]int) bool {
	return input_field == correct_field
}
