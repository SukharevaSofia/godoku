package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	log.SetOutput(os.Stderr)

	var input_val string
	gotta_continue := true
	for gotta_continue {
		fmt.Println(ACTION_CHOISE)
		fmt.Scan(&input_val)

		switch_val, _ := strconv.Atoi(input_val)

		switch switch_val {
		case 1:
      fmt.Println("generating...")
			field := generate()
			fmt.Println("")
			_ = getUserField(field, 27)
			gotta_continue = false

		case 2:
      fmt.Println("generating...")
			field := generate()
			fmt.Println("")
			_ = getUserField(field, 27)
			db_manager()
			fmt.Println("enter name of the save")
			var name string
			fmt.Scan(&name)
			add_field(field, name, false)
			gotta_continue = false

		case 3:
			show_fields()
			gotta_continue = false

		default:
			fmt.Println("incorrect input value")
		}

	}
}
