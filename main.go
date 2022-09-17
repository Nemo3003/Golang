package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

//go mod init
//go run main.go
func main(){
	//Read an input
	reader := bufio.NewReader(os.Stdin)
	
	whatToSay := doctor.Intro()
	
	fmt.Println(whatToSay)

	for {
		fmt.Print("--> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		
		if userInput == "quit"{
			break
		}else{
			fmt.Println(doctor.Response(userInput))
		}

		
	}
	
}
