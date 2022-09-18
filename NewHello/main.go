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
// For: since there is no condition to stop it, it will run forever until someone types "quit"
	for {
		fmt.Print("--> ")
		userInput, _ := reader.ReadString('\n')
		//This is in case the user uses a windows machine
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		//This is for a Linux/Mac OS user
		userInput = strings.Replace(userInput, "\n", "", -1)
		
		if userInput == "quit"{
			break
		}else{
			fmt.Println(doctor.Response(userInput))
		}

	}
	
}
// To compile it with the .exe
// go build -o eliza eliza.exe main.go
// ./eliza