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

