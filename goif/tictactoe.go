package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

type tictacboard [3][3]rune

func main() {
	rand.Seed(time.Now().UnixNano())

	var playerMove bool
	var whoWon string
	var win bool
	var board tictacboard

	fmt.Println("Starting Game: Board is empty")
	board.displayBoard()

	if rand.Intn(2) == 0 {
		playerMove = true
	} else {
		playerMove = false
	}

	for i := 0; i < 9; i++ {
		if playerMove {
			fmt.Println("Player's Move: ", i+1)
			board.playerMove()
			playerMove = false
		} else {
			fmt.Println("Computer Move: ", i+1)
			time.Sleep(time.Second)
			board.computerMove()
			playerMove = true
		}

		if whoWon, win = board.checkWin(); win {
			break
		}

		board.displayBoard()
	}

	fmt.Printf("*****%v won*****\nFinal Board View:\n", whoWon)
	board.displayBoard()
}

func (t *tictacboard) displayBoard() {
	writer := tabwriter.NewWriter(os.Stdout, 1, 4, 1, '\t', tabwriter.AlignRight)
	divider := "-------------------------"
	fmt.Fprintf(writer, divider)
	for i := 0; i < 3; i++ {
		fmt.Fprintf(writer, "\n|")
		for j := 0; j < 3; j++ {
			fmt.Fprintf(writer, " %c\t|", t[i][j])
		}
		fmt.Fprintf(writer, "\n")
		fmt.Fprintf(writer, divider)
	}
	fmt.Fprintf(writer, "\n")
	writer.Flush()
}

func (t *tictacboard) playerMove() {
	var row, col int
	fmt.Print("Enter the Row(1-3) and the Column(1-3) positions: ")
	if _, err := fmt.Scan(&row, &col); err == nil {
		row--
		col--
		if (row >= 0 && row <= 2) && (col >= 0 && col <= 2) && (t[row][col] == 0) {
			t[row][col] = 'X'
		} else {
			fmt.Println("Invalid move, or position not empty. Try again")
			t.playerMove()
		}
	} else {
		fmt.Println("Invalid move, or position not empty. Try again")
		t.playerMove()
	}
}

func (t *tictacboard) computerMove() {
	var row, col int
	for {
		row = rand.Intn(3)
		col = rand.Intn(3)
		if t[row][col] == 0 {
			t[row][col] = 'O'
			break
		}
	}
}

func (t *tictacboard) checkWin() (string, bool) {
	for i := 0; i < 3; i++ {
		if t[i][0] == t[i][1] && t[i][0] == t[i][2] && t[i][0] != 0 {
			if t[i][0] == 'X' {
				return "Player", true
			} else {
				return "Computer", true
			}
		}
		if t[0][i] == t[1][i] && t[0][i] == t[2][i] && t[0][i] != 0 {
			if t[0][i] == 'X' {
				return "Player", true
			} else {
				return "Computer", true
			}
		}
	}
	if t[0][0] == t[1][1] && t[0][0] == t[2][2] && t[0][0] != 0 {
		if t[0][0] == 'X' {
			return "Player", true
		} else {
			return "Computer", true
		}
	}
	if t[0][2] == t[1][1] && t[0][2] == t[2][0] && t[0][2] != 0 {
		if t[0][2] == 'X' {
			return "Player", true
		} else {
			return "Computer", true
		}
	}
	return "", false
}
