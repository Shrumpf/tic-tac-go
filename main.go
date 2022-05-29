package main

import "fmt"

func main() {
	/*
		|x|x|x|
		|-|o|-|
		|-|-|x|
	*/
	won := false
	ls := [9]string{"", "", "", "", "", "", "", "", ""}
	var input int
	currentPlayer := "x"

	for !won {
		fmt.Scan(&input)
		if input >= 0 && input <= 8 {
			ls[input] = currentPlayer

			if currentPlayer == "x" {
				currentPlayer = "o"
			} else {
				currentPlayer = "x"
			}

			winner := checkWin(ls)

			if len(winner) > 0 {
				won = true
				fmt.Println("Gz", winner)
			}

			printBoard(ls)
		}
	}
}

func checkWin(ls [9]string) string {
	checks := [][]string{
		ls[0:3], ls[3:6], ls[6:9],
		{ls[0], ls[3], ls[6]},
		{ls[1], ls[4], ls[7]},
		{ls[2], ls[5], ls[8]},
		{ls[0], ls[4], ls[8]},
		{ls[2], ls[4], ls[6]},
	}

	for i := 0; i < len(checks); i++ {
		if allEqual(checks[i]) {
			return checks[i][0]
		}
	}

	return ""
}

func allEqual(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] || a[i] == "" {
			return false
		}
	}
	return true
}

func printBoard(ls [9]string) {
	fmt.Println("")
	fmt.Println(ls[0:3])
	fmt.Println(ls[3:6])
	fmt.Println(ls[6:9])
}
