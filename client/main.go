package main

import (
	"fmt"
	"math/rand"
	"time"
	//"strconv"
)


func main() {

	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 4
	cp := rand.Intn(max-min) + min

	var user1 string
	var ch1 int


	fmt.Println("Welcome to Rock Paper Scissor \n")
	
	fmt.Println("User enter your name. ")
	fmt.Scanln(&user1)	
	//fmt.Println("User 2 enter your name. ")
	//fmt.Scanln(&user2)

	fmt.Println("\n" + user1 + " v/s " + "Computer")
	
	fmt.Println("\n" + user1 + " Enter your choice - 1. Rock; 2. Paper; 3. Scissor")
	fmt.Scanln(&ch1)

	//fmt.Println("\n" + user2 + " Enter your choice - 1. Rock; 2. Paper; 3. Scissor")
	//fmt.Scanln(&ch2)

	if(ch1 == 1 && cp == 1) {
		fmt.Println("\n Game is a draw")
	}
	
	if(ch1 == 1 && cp == 2) {
		fmt.Println("\n User 2 is a winner")
	}

	if(ch1 == 1 && cp == 3) {
		fmt.Println("\n User 1 is a winner")
	}

	if(ch1 == 2 && cp == 1) {
		fmt.Println("\n User 1 is a winner")
	}
	
	if(ch1 == 2 && cp == 2) {
		fmt.Println("\n Game is a draw")
	}

	if(ch1 == 2 && cp == 3) {
		fmt.Println("\n User 2 is a winner")
	}

	if(ch1 == 3 && cp == 1) {
		fmt.Println("\n User 2 is a winner")
	}
	
	if(ch1 == 3 && cp == 2) {
		fmt.Println("\n User 1 is a winner")
	}

	if(ch1 == 3 && cp == 3) {
		fmt.Println("\n Game is a draw")
	}
}
