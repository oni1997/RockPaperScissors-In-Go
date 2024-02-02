package main
// always start a file with package main

import(
	"fmt" // import fmt for printing to the console
	"math/rand" //the equivalent of import random.randint in python to generate  random numbers
)

//consts can never be changed  after they are declared, unlike variables which can change value
const(
	Rock = 0
	Paper = 1
	Scissor = 2
)

//var are subject to  change and can hold values of different types
var actions = map[int] string{
	Rock: "Rock",
	Paper: "Paper",
	Scissor: "Scissors",
}

//func is our function that  takes no arguments but returns nothing (void). It has access to all other functions and variables within this
func main(){
	// Println prints out the specified values and automatically appends a newline character at the end.	fmt.Println("Welcome to My Rock Paper Scissors")
	fmt.Println("Like it says you have 3 Choices Rock Paper Scissors")
	fmt.Println("Rock = 0\nPaper = 1\nScissors = 2")

	var userInput int
	for {
		fmt.Print("Input")
		_, err := fmt.Scan(&userInput)
		if err == nil  && (userInput >= Rock && userInput <= Scissor){
			break
		}else{
			fmt.Println("Invalid Input. Please enter either 0 or 1 or 2")
		}

	}
	computerInput := rand.Intn(3) + 1 // rand.Intn returns a random number between 0 and n-1 so we add 1 to get 

	var result string
	if userInput == Rock && computerInput == Scissor{
		result = "Congratulations! You Win!"
	}

	computerAction := actions[computerInput]
	userAction := actions[userInput]
	fmt.Println("%vUser's action was %vComputer's action was %v", result, userAction, computerAction)
}