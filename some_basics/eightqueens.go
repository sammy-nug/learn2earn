package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int

	var solve func(row int)
	solve = func(row int) {
		if row == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '1'))
			}
			z01.PrintRune('\n')
			return
		}

		for col := 0; col < 8; col++ {
			safe := true

			for i := 0; i < row; i++ {
				// same column
				if board[i] == col ||
					// diagonal conflict
					board[i]-col == row-i ||
					col-board[i] == row-i {

					safe = false
					break
				}
			}

			if safe {
				board[row] = col
				solve(row + 1)
			}
		}
	}

	solve(0)
}
