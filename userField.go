package main

import (
	"math/rand"
)

func getUserField(fullField []int, hints int) []int {

	hintsIndexes := getIndexesToHide(hints)

	userField := make([]int, 81)
	for i := range fullField {

		//means the index is that of a hint's
		if hintsIndexes[i] {
			userField[i] = fullField[i]
		} else {
			userField[i] = 0
		}
	}
	printField(userField)
	return userField
}

func getIndexesToHide(hints int) map[int]bool {
	randomIndexes := rand.Perm(81)
	indexesMap := make(map[int]bool)
	for i := 0; i < hints; i++ {
		index := randomIndexes[i]
		indexesMap[index] = true
	}
	return indexesMap
}
