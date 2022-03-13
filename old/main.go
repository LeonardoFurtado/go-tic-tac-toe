package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const playerOne = "X"
const playerTwo = "O"

var board = [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func showBoard() {
	// Print the board
	CallClear()
	for i := 0; i < 3; i++ {
		fmt.Printf("%s | %s | %s\n", board[i][0], board[i][1], board[i][2])
		if i < 2 {
			fmt.Printf("---------\n")
		}
	}
}

func checkWinner(player string) bool {
	// Lines
	if board[0][0] == player && board[0][1] == player && board[0][2] == player {
		return true
	}
	if board[1][0] == player && board[1][1] == player && board[1][2] == player {
		return true
	}
	if board[2][0] == player && board[2][1] == player && board[2][2] == player {
		return true
	}

	//Columns
	if board[0][0] == player && board[1][0] == player && board[2][0] == player {
		return true
	}

	if board[0][1] == player && board[1][1] == player && board[2][1] == player {
		return true
	}

	if board[0][2] == player && board[1][2] == player && board[2][2] == player {
		return true
	}

	//Diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

func changePlayer(player string) string {
	if player == playerOne {
		return playerTwo
	}
	return playerOne
}

func checkValidPosition(x, y int) bool {
	result := board[x][y] == " " && x <= 2 && x >= 0 && y <= 2 && y >= 0
	return result
}

func insertOnBoard(x, y int, player string) {
	validPosition := checkValidPosition(x, y)

	for !validPosition {
		fmt.Println("This position is not valid, please choice another position")
		_, _ = fmt.Scanf("%d %d", &x, &y)
		validPosition = checkValidPosition(x, y)
	}

	fmt.Printf("Value inserted at position %d %d\n", x, y)
	time.Sleep(2 * time.Second)
	board[x][y] = player
}

func makeIteractions() {
	anyWin := false
	currentPlayer := playerOne
	var positionX, positionY int
	for !anyWin {
		fmt.Printf("%s choice a position:\n", currentPlayer)
		_, _ = fmt.Scanf("%d %d", &positionX, &positionY)
		insertOnBoard(positionX, positionY, currentPlayer)
		anyWin = checkWinner(currentPlayer)
		if anyWin {
			fmt.Printf("%s wins! Press any key to exit.\n", currentPlayer)
			_, _ = fmt.Scanf("%s", currentPlayer)
			return
		}
		currentPlayer = changePlayer(currentPlayer)
		showBoard()
	}
}

func main() {
	showBoard()
	fmt.Println("Game Started!!!")
	time.Sleep(2 * time.Second)
	makeIteractions()
	fmt.Println("Ending game...")
}
