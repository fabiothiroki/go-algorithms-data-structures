package sudoku

import (
	"strconv"
)

// IsValidSudoku determines if a sudoku is valid
func IsValidSudoku(board [][]string) bool {

	for i := range board {
		for j := range board {

			if isEmpty(board[i][j]) {
				continue
			}

			isPositionValid := validatePosition(board, i, j)

			if isPositionValid == false {
				return false
			}

		}
	}

	return true
}

func validatePosition(board [][]string, line int, column int) bool {

	value := board[line][column]

	for i := range board {
		if board[line][i] == value && i != column {
			return false
		}

		if board[i][column] == value && i != line {
			return false
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
