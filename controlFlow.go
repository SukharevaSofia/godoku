package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)
func clear(){
  fmt.Println(CLEAN_TERMINAL)
}

func control_flow() {
  clear()
	var input_val string

	// generating or printing fields
	for {
		fmt.Print(GEN_OR_SHOW)
		fmt.Scan(&input_val)
		switch_val, _ := strconv.Atoi(input_val)

		switch switch_val {
    // user decided to gen field
		case 1:
      hints := choose_difficulty()
			fmt.Println("generating...")
			field := generate()
			fmt.Println("Generahion completed!")
			play_or_not(field, hints)

    //user decided to show fields
		case 2:
      is_obscured := choose_showing()
			show_fields(is_obscured)
			log.Println("terminating after showing fields")
			os.Exit(0)

		default:
			fmt.Println(BAD_INPUT)
		}
	}
}

func choose_showing() bool{
  clear()
	var tmp string
	for {
		fmt.Print(SHOW_FULL_OR_HINT)
		fmt.Scan(&tmp)
		is_obsc, _ := strconv.Atoi(tmp)
		switch is_obsc {
    //show fields with hidden values
		case 1:
			return true
      //show fields as is without hiding
		case 2:
			return false
		default:
			fmt.Println(BAD_INPUT)
		}
	}
}

func choose_difficulty() int {
  clear()
	var tmp string
	for {
		fmt.Print(DIFFICULTY_CHOICE)
		fmt.Scan(&tmp)
		difficulty, _ := strconv.Atoi(tmp)
		switch difficulty {
    //easy
		case 1:
			return 60
    //medium
		case 2:
			return 40
    //hard
		case 3:
			return 27
		default:
			fmt.Println(BAD_INPUT)
		}
	}
}

func play_or_not(full_field [81]int, hints int) {
  clear()
  hints_loc, play_field := obscure_field(full_field, hints)
  print_field(play_field)

	var input_val string
	gotta_continue := true
	for gotta_continue {
		fmt.Print(PLAY_CHOICE)
		fmt.Scan(&input_val)
		switch_val, _ := strconv.Atoi(input_val)

		switch switch_val {
    //user decided to play
		case 1:
		  fmt.Println("STAB! NO GAMEPLAY YET!")
      os.Exit(0)
    //user decided not to play
		case 2:
      save_or_not(full_field, play_field, hints_loc)
		default:
			fmt.Println(BAD_INPUT)
		}
	}
}

func save_or_not(full_field, play_field, hints_loc [81]int) {
  clear()
	var input_val string

	for {
    print_field(play_field)
		fmt.Print(SAVE_CHOICE)
		fmt.Scan(&input_val)
		switch_val, _ := strconv.Atoi(input_val)
		switch switch_val {
    // to save the field
		case 1:
			db_manager()
      fmt.Print("\nenter name of the save:\n > ")
			var name string
			fmt.Scan(&name)
      full_id, play_id, _ := save_field(full_field, play_field, hints_loc, name, false)
      fmt.Printf("\nID of the saved field = %d. To see the answer look for field # %d\n", play_id, full_id)
			log.Println("terminating the programm after saving")
			os.Exit(0)
    //abandon the generated field, for it to be forever gone in the flow of time...
		case 2:
			log.Println("terminating the programm, no saving")
			os.Exit(0)
		default:
			fmt.Println(BAD_INPUT)
		}
	}
}

// returns the number of hints
func gameplay(full_field, play_field []int, hints_locs map[int]bool){}

func readUserInput() {
	// TODO implement reading from the buttons
}
