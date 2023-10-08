package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz")

	//Enter your name scan+ print
	var name string

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)                                      //& is for allowing to store the value anywhere in RAM
	fmt.Printf("Hello, %v, welcome to the game\n", name) // "\n" is break

	//enter age + condition + scan+ print
	var age uint

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age) //& is for allowing to store the value anywhere in RAM
	fmt.Println(age >= 10)

	if age >= 10 {
		fmt.Println("You are old enoguh to play")
	} else {
		fmt.Println("You are not old enough to play, sorry")
		return // exit the main function
	}

	score := 0
	num_questions := 2

	//first question if, else statements + scan+ print
	fmt.Printf("What's the current year? ")
	var answer int
	fmt.Scan(&answer)

	if answer == 2023 {
		fmt.Println("Good, you are correct")
		score++
	} else {
		fmt.Println("Nah, incorrect")
	}

	//second question if, else statements + scan + print
	fmt.Printf("How many days does a year have? ")
	var answer2 int
	fmt.Scan(&answer2)

	if answer2 == 365 {
		fmt.Println("Good, you are correct")
		score++
	} else {
		fmt.Println("Nah, incorrect")
	}

	// print score
	fmt.Printf("You scored %v out if %v. \n", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100 //type conversion necessary
	fmt.Printf("You scored: %v%%.", percent)                   // %% needed to print % to console
}
