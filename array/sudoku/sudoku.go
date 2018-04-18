package sudoku

import (
	"strconv"
)

// IsValidSudoku determines if a sudoku is valid
func IsValidSudoku(board [][]string) bool {

	var validLines [9]bool
	var validColumns [9]bool
	var validQuadrants [3][3]bool

	for i := range board {
		for j := range board {

			if isEmpty(board[i][j]) {
				continue
			}

			if !validLines[i] {
				if !validateLine(board, i, j) {
					return false
				}
				validLines[i] = true
			}

			if !validColumns[j] {
				if !validateColumn(board, i, j) {
					return false
				}
				validColumns[j] = true
			}

			if !validQuadrants[i/3][j/3] {
				if !validateQuadrant(board, i/3, j/3) {
					return false
				}

				validQuadrants[i/3][j/3] = true
			}
		}
	}

	return true
}

func validateLine(board [][]string, line int, column int) bool {

	m := make(map[string]bool)

	for i := range board {
		if isEmpty(board[line][i]) {
			continue
		}

		value := board[line][i]

		if m[value] == true {
			return false
		}

		m[value] = true
	}

	return true
}

func validateColumn(board [][]string, line int, column int) bool {

	m := make(map[string]bool)

	for i := range board {
		if isEmpty(board[i][column]) {
			continue
		}

		value := board[i][column]

		if m[value] == true {
			return false
		}

		m[value] = true
	}

	return true
}

func validateQuadrant(board [][]string, lineQuadrant int, columnQuadrant int) bool {
	var minLine int
	var minColumn int
	if lineQuadrant == 0 {
		minLine = 0
	} else if lineQuadrant == 1 {
		minLine = 3
	} else {
		minLine = 6
	}

	if columnQuadrant == 0 {
		minColumn = 0
	} else if columnQuadrant == 1 {
		minColumn = 3
	} else {
		minColumn = 6
	}

	m := make(map[string]bool)

	for i := minLine; i <= minLine+2; i++ {
		for j := minColumn; j <= minColumn+2; j++ {
			if isEmpty(board[i][j]) {
				continue
			}

			value := board[i][j]

			if m[value] == true {
				return false
			}

			m[value] = true
		}
	}
	return true
}

func isEmpty(boardPosition string) bool {
	if _, err := strconv.Atoi(boardPosition); err == nil {
		return false
	}
	return true
}
