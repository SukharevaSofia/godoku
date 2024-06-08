package main

import (
	"fmt"
	"math/rand"
)

func getNeighboursMap(field, row []int, index int) map[int]bool {
	neighbours := make(map[int]bool, 10)

	var value int

	for i := range row {
		value = row[i]
		neighbours[value] = true
	}

	column := index % 9

	for j := column; j < index; j += 9 {
		value = field[j]
		neighbours[value] = true
	}
	return neighbours
}

func getValidNumber(field, unfinishedRow []int, curCell, curFieldRow, counterOfIterations int)int{
  var num int
	neighbours := getNeighboursMap(field, unfinishedRow, curCell)
  for{
		num = rand.Intn(9) + 1

		if neighbours[num] {
			continue
		}

		if !(checkValidZone(field, curFieldRow, curCell%9, num)) {
			fmt.Println("unfinished row", unfinishedRow, "cucurFieldRow", curFieldRow, "culcurCell%9", curCell%9, "num", num, "nuigh", neighbours)
			continue
		}
    break
  }
  return num
}

func generate() []int {
	field := generateFirstThreeRows()
	curFieldRow := 27

	for curFieldRow < 81 {
		randRow := make([]int, 9)
		for curCell := curFieldRow; curCell < curFieldRow+9; curCell++ {
      conterOfIterations := 0
      num := getValidNumber(field, randRow, curCell, curFieldRow, counterOfIterations)
			randRow[curCell%9] = num
		}

		for row := range randRow {
			field[curFieldRow+row] = randRow[row]
		}

		curFieldRow += 9
	}
	printField(field)
	return field
}

func generateFirstThreeRows() []int {

	field := make([]int, 81)
	randRow := make([]int, 9)

	for curFieldRow := 0; curFieldRow < 27; curFieldRow += 9 {
		for {
			isBad := false
			randRow = getRandomRow()
			for curFieldColumn := range randRow {
				if !checkValid(field, curFieldRow, curFieldColumn, randRow[curFieldColumn]) {
					isBad = true
					break
				}
			}

			if isBad {
				continue
			}

			for column := range randRow {
				field[curFieldRow+column] = randRow[column]
			}

			break
		}
	}
	printField(field)
	return field
}

func checkValidRow(row []int, value int) bool {
	for i := range row {
		if row[i] == value {
			return false
		}
	}
	return true
}

func checkValidColumn(field []int, column, value int) bool {

	for curCell := column; curCell < 81; curCell += 9 {
		if field[curCell] == value {
			return false
		}
	}

	return true
}

func checkValidZone(field []int, row, column, value int) bool {

	cellFirstRow, cellFirstCoulumn := getFirstInZone(row, column)
	currCell := cellFirstRow + cellFirstCoulumn
	lastZoneCell := (cellFirstRow + 2*9) + (cellFirstCoulumn + 2)
	for currCell < (lastZoneCell) {

		if currCell >= 81 {
			break
		}

		for j := 0; j < 3; j++ {
			cellInColumn := currCell + j
			if field[cellInColumn] == value {
				fmt.Println("поле от", cellInColumn, "=", field[cellInColumn], "значение", value)
				return false
			}
		}

		currCell += 9
	}
	return true
}

func checkValid(field []int, row, column, value int) bool {
	return checkValidColumn(field, column, value) && checkValidZone(field, row, column, value)
}

func getRandomRow() []int {
	randomRowContents := rand.Perm(9)

	// sudoku goes from 1 to 9 therefore we increment here
	for i := range randomRowContents {
		randomRowContents[i]++
	}

	return randomRowContents
}

func getFirstInZone(row, column int) (int, int) {
	if row <= 18 {
		row = 0
	} else if row <= 45 {
		row = 27
	} else {
		row = 54
	}
	if column <= 2 {
		column = 0
	} else if column <= 5 {
		column = 3
	} else {
		column = 6
	}
	return row, column
}
