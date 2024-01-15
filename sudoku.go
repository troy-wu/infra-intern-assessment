package main


import (
	"fmt"
	"reflect"
	"testing"
)

type SudokuSolver struct {
	board [][]int
}

func NewSudokuSolver(board [][]int) *SudokuSolver {
	return &SudokuSolver{board: board}
}

func (s *SudokuSolver) SolveSudoku() [][]int {
	s.solve()
	return s.board
}

func (s *SudokuSolver) solve() bool {
	empty := findEmptyCell(s.board)
	if empty == nil {
		return true // All cells are filled
	}

	row, col := empty[0], empty[1]

	for num := 1; num <= 9; num++ {
		if isValid(s.board, row, col, num) {
			s.board[row][col] = num

			if s.solve() {
				return true
			}

			s.board[row][col] = 0 // Backtrack if the current num doesn't lead to a solution
		}
	}

	return false
}

func findEmptyCell(board [][]int) []int {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func isValid(board [][]int, row, col, num int) bool {
	return !usedInRow(board, row, num) &&
		!usedInCol(board, col, num) &&
		!usedInBox(board, row-row%3, col-col%3, num)
}

func usedInRow(board [][]int, row, num int) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return true
		}
	}
	return false
}

func usedInCol(board [][]int, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(board [][]int, startRow, startCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return true
			}
		}
	}
	return false
}
