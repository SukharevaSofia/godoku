package main

import (
	"fmt"
)

const (
	CHOOSE_DIFFICULTY = "Please, choose the level of difficulty:\n" +
		"1) Easy\n" +
		"2) Medium\n" +
		"3) Hard\n"
)

type field struct {
	Id        int64
	Save_name string
	Is_solved bool
}

func printField(field []int) {
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
