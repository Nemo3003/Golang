package main

import (
	"bufio"	
	"fmt"
	"os"
)

const stringRepeat = "and press enter once you are done."

func main() {
	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

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

	answer = firstNumber * secondNumber - subtraction
	fmt.Println("The answer is", answer)
}