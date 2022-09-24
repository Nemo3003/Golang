package main

import (
	"fmt"
)

func main(){
	// Init game board with empty strings

	//Board := [3][3] string{}

	// var to carry current player name
	player := "x"

	fmt.Println("player",player)

	// read row value

	var row int
	fmt.Println("Please enter number 0,1 or 2")
	fmt.Scanln(&row)

}