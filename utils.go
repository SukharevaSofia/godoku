package main

import "fmt"

const (
	CHOOSE_DIFFICULTY = "Please, choose the level of difficulty:\n" +
		"1) Easy\n" +
		"2) Medium\n" +
		"3) Hard\n"
)

func printField(field []int) {
	for i, el := range field {
		
    if i%9 == 0 {
			fmt.Printf("\n")
		}

    if (i % 3 == 0) && ((i % 9) != 0) {
      fmt.Printf("|")
    }
    if el == 0{
		  fmt.Printf(" · ")
    }else{
		  fmt.Printf(" %d ", el)
    }
    if (i == 26) || (i == 53){
      fmt.Printf("\n---------+---------+--------")
    }
	}
}
