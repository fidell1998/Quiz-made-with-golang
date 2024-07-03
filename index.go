package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Welcome to my quiz made with Golang!")

	fmt.Printf("Enter your name: ")
	var usersName string
	fmt.Scan(&usersName)

	fmt.Printf("Hello %v! ", usersName)

	fmt.Printf("How old are you?\n Enter your age: ")
	var usersAge uint
	fmt.Scan(&usersAge)

	usersScore := 0

	if usersAge > 0 {
		fmt.Printf(" %v years old, pretty nice, now let's begin.", usersAge)
	} else {
		fmt.Println("We can not continue if you are going to put a negative age or age of zero.")
		return
	}

	fmt.Printf("\nQuestion 1: What is better, the RTX 3080 or RTX 3090? ")
	var ans1Ques1 string
	var ans2Ques1 string
	fmt.Scan(&ans1Ques1, &ans2Ques1)

	ans1Ques1 = strings.ToUpper(ans1Ques1)
	ans2Ques1 = strings.ToUpper(ans2Ques1)

	if ans1Ques1 == "RTX" && ans2Ques1 == "3090" {
		fmt.Println("Correct!")
		usersScore += 1
	} else {
		fmt.Println("Incorrect, sorry")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? 4, 8, or 12? ")
	var ansQues2 uint
	fmt.Scan(&ansQues2)

	if ansQues2 == 12 {
		fmt.Println("Nice! Correct")
		usersScore += 1
	} else {
		fmt.Println("Oops, incorrect")
	}

	fmt.Printf("True or false? www. stands for world wide website. ")
	var ansQues3 bool
	fmt.Scan(&ansQues3)

	if ansQues3 == false {
		fmt.Printf("Correct, the statement was false.")
		usersScore += 1
	} else {
		fmt.Printf("incorrect, the statement was was.")
	}

	fmt.Printf("\nfinal question, How many bits are in a byte? 1, 8, or 16? ")
	var ansQues4 uint
	fmt.Scan(&ansQues4)

	if ansQues4 == 8 {
		fmt.Println("Nice, correct")
		usersScore += 1
	} else {
		fmt.Println("Incorrect, good try though")
	}

	fmt.Printf("You scored %v out 4 correct! Thanks for playing!\n", usersScore)
	percent := (float64(usersScore) / float64(4)) * 100
	fmt.Printf("You scored %v%%! Thanks for playing.", percent)
}
