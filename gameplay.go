package main

import (
	"fmt"
	"os"
	"os/exec"
)

func game_controller(full_field, play_field, hints_loc [81]int) {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b []byte = make([]byte, 3)

	var arr_first_symb byte = 27
	var arr_second_symb byte = 91
	var arr_top byte = 65
	var arr_bot byte = 66
	var arr_right byte = 67
	var arr_left byte = 68

	local_field := play_field
	cursor_index := 0
	clear()
	print_field(local_field)
	for {
		clear()
		print_field_with_coursor(local_field, cursor_index)
		fmt.Print(INSTRUCTIONS)
		os.Stdin.Read(b)
		// pressed one of the arrows
		switch b[0] {
		case arr_first_symb:
			if b[1] == arr_second_symb {
				switch b[2] {
				case arr_top:
					cursor_index = cursor_up(cursor_index, local_field)
				case arr_bot:
					cursor_index = cursor_down(cursor_index, local_field)
				case arr_right:
					cursor_index = cursor_right(cursor_index, local_field)
				case arr_left:
					cursor_index = cursor_left(cursor_index, local_field)
				default:
					fmt.Println("weird input")
				}
			} else {
				fmt.Println("weird imput")
			}

		// pressed 'enter'
		case 10:
			fmt.Printf("Please, input a number\n> ")
			os.Stdin.Read(b)
			if 48 <= b[0] && b[0] <= 57 {
				value_to_insert := int(b[0]) - 48
				local_field = write_into_cell(cursor_index, value_to_insert, local_field)
				if equal_fields(local_field, full_field) {
					we_have_a_winner(local_field, full_field, hints_loc)
				} else {
					clear()
					print_field_with_coursor(local_field, cursor_index)
				}
			} else {
				bad_input()
			}
		// pressed escape
		case 99:
			exec.Command("stty", "-F", "/dev/tty", "echo").Run()
			save_or_not(full_field, local_field, hints_loc, false)
		default:
			bad_input()
		}

	}
}

func write_into_cell(cursor_index, new_value int, field [81]int) [81]int {
	field[cursor_index] = new_value
	return field
}

func cursor_up(cursor_index int, field [81]int) int {
	clear()
	if cursor_index <= 8 {
		cursor_index += 72
	} else {
		cursor_index -= 9
	}
	print_field_with_coursor(field, cursor_index)
	return cursor_index
}

func cursor_down(cursor_index int, field [81]int) int {
	clear()
	if cursor_index >= 72 {
		cursor_index = cursor_index % 9
	} else {
		cursor_index += 9
	}
	print_field_with_coursor(field, cursor_index)
	return cursor_index
}

func cursor_right(cursor_index int, field [81]int) int {
	clear()
	if cursor_index%9 == 8 {
		cursor_index -= 8
	} else {
		cursor_index += 1
	}
	print_field_with_coursor(field, cursor_index)
	return cursor_index
}

func cursor_left(cursor_index int, field [81]int) int {
	clear()
	if cursor_index%9 == 0 {
		cursor_index += 8
	} else {
		cursor_index -= 1
	}
	print_field_with_coursor(field, cursor_index)
	return cursor_index
}

func we_have_a_winner(full_field, local_field, hints_loc [81]int) {
	clear()
	print_field(full_field)
	fmt.Printf("\n \033[93m %s \033[m", WINNER)
	fmt.Println("\nPress 'enter' to continue.")
	var b []byte = make([]byte, 3)
  for{
	  os.Stdin.Read(b)
    if b[0] == 10{
      break
    }
  }
  exec.Command("stty", "-F", "/dev/tty", "echo").Run()
	save_or_not(full_field, local_field, hints_loc, true)
}
