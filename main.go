package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func solveSudoku(board [][]byte) {
	available := availablePos(board)
	if backTrack(board, available, 0) { // sudokuH
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func availablePos(board [][]byte) [][]int {
	var ava [][]int
	for y, row := range board {
		for x, e := range row {
			if e == '.' {
				ava = append(ava, []int{x, y})
			}
		}
	}
	return ava
}

func backTrack(tracker [][]byte, available [][]int, count int) bool {
	if count == 81 {
		return true
	}
	x := count / 9
	y := count % 9
	if tracker[x][y] != '.' {
		return backTrack(tracker, available, count+1)
	}

	var i byte
	for i = 49; i <= 57; i++ {
		if numNotPresentInRow(tracker[x], i) && numNotPresentInColumn(tracker, y, i) &&
			numNotPresentInSquare(tracker, x, y, i) && tracker[x][y] == '.' {

			tracker[x][y] = i
			if backTrack(tracker, available, count+1) {
				return true
			}
			tracker[x][y] = '.'
		}
	}
	return false
}

//These are helpers to check if the number is present in row, column or 3X3 square region
func numNotPresentInRow(trackerRow []byte, i byte) bool {
	for _, val := range trackerRow {
		if i == val {
			return false
		}
	}
	return true
}

func numNotPresentInColumn(tracker [][]byte, y int, i byte) bool {
	for _, row := range tracker {
		if row[y] == i {
			return false
		}
	}
	return true
}

func numNotPresentInSquare(tracker [][]byte, x, y int, i byte) bool {
	a1 := 0
	a2 := 0

	b1 := 0
	b2 := 0
	if x >= 0 && x < 3 {
		a1 = 0
		a2 = 3
	} else if x >= 3 && x < 6 {
		a1 = 3
		a2 = 6
	} else {
		a1 = 6
		a2 = 9
	}

	if y >= 0 && y < 3 {
		b1 = 0
		b2 = 3
	} else if y >= 3 && y < 6 {
		b1 = 3
		b2 = 6
	} else {
		b1 = 6
		b2 = 9
	}
	for m := a1; m < a2; m++ {
		for n := b1; n < b2; n++ {
			if tracker[m][n] == i {
				return false
			}
		}
	}
	return true
}

func printBoard(board [][]byte) {
	for _, row := range board {
		for i, e := range row {
			z01.PrintRune(rune(e))
			if i != len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func validW(param []string) bool {

	lParams := len(param)
	valid := false

	if lParams < 9 || lParams > 9 {
		valid = false
	} else if lParams == 9 {
		for i := 0; i < lParams; i++ {
			arg := param[i]
			lParam := len(arg)
			//fmt.Println("ARG = ", arg)
			// ============== validaion for len
			if lParam < 9 || lParam > 9 {
				valid = false
				//fmt.Println("Error")
				break
			} else {

				valid = true
			}
			if valid == true {
				for j := 0; j < lParam; j++ {

					indexArg := arg[j]
					if (indexArg == 46 || indexArg > 47) && (indexArg == 46 || indexArg < 58) { // 46 == . , 48 - 0, 57 - 9
						//fmt.Println(arg[j])
						valid = true
					} else {
						valid = false
						//fmt.Println("Error")
						break
					}

				}
			}

		}
	}
	return valid
}

func main() {
	//".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
	param1 := os.Args[1:]

	valid := validW(param1)
	var board [][]byte

	for _, v := range os.Args[1:] {
		board = append(board, []byte(v))
	}

	if valid == true {

		solveSudoku(board)

	} else {
		fmt.Println("Error")
	}
}
