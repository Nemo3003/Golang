package main

//basic types (numbers, strings, booleans)

//agregate types (array, struct)

//reference types (pointers, slices, maps, functions, channels)

//interface type

import (
	"fmt"
	"github.com/eiannone/keyboard"
)
type Car struct{
	numberOfTires int
	Luxury bool
	Bucketseats bool
	Make string
	Model string
	Year int
}

var keyPressChan chan rune
func main() {
	/*var myString [3] string
	myString[0] = "cat"
	myString[1] = "cat1"
	myString[2] = "cat2"
	fmt.Printf("%s\n", myString[0])*/
	//--------------STRUCT
	myCar := Car{
		numberOfTires: 4,
		Luxury: false,
		Bucketseats: true,
		Model: "XC",
		Make: "Volkswagen",
		Year: 2019,
	}
	fmt.Printf("My car is a %d %s %s\n", myCar.Year, myCar.Make, myCar.Model)
	main2()

}

func main2(){
	/*x := 10

	myFirstPointer := &x

	
  POINTERS
	fmt.Println("x is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValuePointer(&x)

	fmt.Println("x is now (after its change)", x)*/
/*
--------------SLICES----------------------
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	animals = append(animals, "fish")

	fmt.Println(animals)

	for i, x := range animals{
		fmt.Println(i,x)
	}*/

	/*
	--------------MAP-------------
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5
	intMap["six"] = 6

	delete(intMap, "one")*/
	

/*func changeValuePointer(num *int){
	*num = 25}
*/

keyPressChan = make(chan rune)

go listenForKeyPress()

fmt.Println("Press any key or q to quit")

_ = keyboard.Open()
defer func(){
	keyboard.Close()
}()

for {
	char, _, _ := keyboard.GetSingleKey()

	if char == 'q' || char == 'Q'{
		break
	}

	keyPressChan <- char 
}
}




func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}