package main

import (
	"math/rand"
)

func obscure_field(full_field [81]int, hints int) ([81]int, [81]int) {

	hints_indexes := get_hints_indexes(hints)

	var obscured_field [81]int
	for i := range full_field {

		//means the index is that of a hint's
		if hints_indexes[i] != 0 {
			obscured_field[i] = full_field[i]
		} else {
			obscured_field[i] = 0
		}
	}
	return hints_indexes, obscured_field
}

func get_hints_indexes(hints int) [81]int {
	random_indexes := rand.Perm(81)
	var indexes_map [81]int
	for i := 0; i < hints; i++ {
		index := random_indexes[i]
		indexes_map[index] = 1
	}
	return indexes_map
}
