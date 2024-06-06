package main

import (
	"fmt"
	"math/rand"
)


func generateFirstThreeRows() []int{ 
  
  field := make([]int, 81)
  randRow := make([]int, 81)

	for curFieldRow := 0; curFieldRow < 27; curFieldRow += 9 {
		for {
      isBad := false
			randRow = getRandomRow()
			for curFieldColumn := range randRow {
				//if the value doesnt fit, we null the whole row and repeat the generation
				if !checkValid(field, curFieldRow/9, curFieldColumn, randRow[curFieldColumn]) {
          isBad = true
					break
				}
			}

      if isBad{
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

func generate() []int {
  field := generateFirstThreeRows()
  randRow := make([]int, 81)

	for curFieldRow := 0; curFieldRow < 81; curFieldRow += 9 {
		for {
      isBad := false
			randRow = getRandomRow()
			for curFieldColumn := range randRow {
				//if the value doesnt fit, we null the whole row and repeat the generation
				if !checkValid(field, curFieldRow/9, curFieldColumn, randRow[curFieldColumn]) {
          isBad = true
					break
				}
			}

      if isBad{
        continue
      }

      for column := range randRow {
        field[curFieldRow+column] = randRow[column]
      }

			break
		}
    fmt.Println("finished row", curFieldRow/9)
	  printField(field)
    counter = 0
  }
	printField(field)
  return field
}

func checkValidColumn(field []int, column, value int) bool{

	for curCell := column; curCell < 81; curCell += 9 {
		if (field[curCell] == value) {
      counter++
			return false
    }
	}

  return true
}

func checkValidZone(field []int, row, column, value int) bool {
	
  cellFirstRow, cellFirstCoulumn := getFirstInZone(row, column)
	currCell := cellFirstRow + cellFirstCoulumn

	for currCell < (getIndex(cellFirstRow + 2, cellFirstCoulumn + 2)) {
    
    if currCell >= 81{
      break
    }

		for j := 0; j < 3; j++ {
      cellInColumn := currCell + j
			if (field[cellInColumn] == value) {
        counter++
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

func getIndex(row, column int) int{
  return row*9 + column
}

func getFirstInZone(row, column int)(int, int){
  return row - (row % 3), column - (column % 3) 
}

