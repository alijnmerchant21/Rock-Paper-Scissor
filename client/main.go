package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	cp := rand.Intn(max-min) + min

	var user1 string
	var ch1 int

	fmt.Println("\nWelcome to Rock Paper Scissor")
	fmt.Println("\nUser enter your name." + "\n")
	fmt.Scanln(&user1)

	fmt.Println("\n" + user1 + " v/s " + "Computer")
	fmt.Println("\n" + user1 + " Enter your choice \n1. Rock \n2. Paper \n3. Scissor\n")
	fmt.Scanln(&ch1)

	if cp == 1 {
		fmt.Println("\n\nComputer's choice is Rock")
	}

	if cp == 2 {
		fmt.Println("\n\nComputer's choice is Paper")
	}

	if cp == 3 {
		fmt.Println("\n\nComputer's choice is Scissor")
	}

	switch ch1 {

	case 1:
		if cp == 1 {
			fmt.Println("\n\n Game is a draw")
		}
		if cp == 2 {
			fmt.Println("\n\n Computer is a winner")
		}
		if cp == 3 {
			fmt.Println("\n\n" + user1 + " is a winner")
		}

	case 2:
		if cp == 1 {
			fmt.Println("\n\n" + user1 + " is a winner")
		}
		if cp == 2 {
			fmt.Println("\n\n Game is a draw")
		}
		if cp == 3 {
			fmt.Println("\n\n Computer is a winner")
		}

	case 3:
		if cp == 1 {
			fmt.Println("\n\n Computer is a winner")
		}
		if cp == 2 {
			fmt.Println("\n\n" + user1 + " is a winner")
		}
		if cp == 3 {
			fmt.Println("\n\n Game is a draw")
		}

	}
}
