package main

import (
	"fmt"
	"myapp/packageone"
	)

func main() {
	//var one = "One"

	newString := packageone.PublicVar
	fmt.Println(newString)

	
}