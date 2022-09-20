package main

import (
	"bufio"	
	"fmt"
	"os"
	"math/rand"
	"time"
)

const stringRepeat = "and press enter once you are done."

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())
	var firstNumber = rand.Intn(8)
	var secondNumber = rand.Intn(8)
	var subtraction = rand.Intn(8)
	answer := firstNumber * secondNumber - subtraction
	
	//---------------------------------------------------
	numbRand( answer, firstNumber, secondNumber, subtraction )
}

func numbRand( answer, firstNumber, secondNumber, subtraction int){
	reader := bufio.NewReader(os.Stdin)
	//display a welcome/instructions\
	fmt.Println("Guess the number game")
	fmt.Println("----------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", stringRepeat)
	reader.ReadString('\n')
	//Take them through the game
	fmt.Println("Multiply your number by", firstNumber, stringRepeat)
	reader.ReadString('\n')

	fmt.Println("Now the result by", secondNumber, stringRepeat)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the name you originally thought of", stringRepeat)
	reader.ReadString('\n')

	fmt.Println("Now subtract:", subtraction, stringRepeat)
	//give them the answer

	
	fmt.Println("The answer is", answer)
}