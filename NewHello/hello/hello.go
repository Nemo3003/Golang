package main

import "fmt"
//go run hello.go 

func main() {
    /*var x int
    var y int

    x = 1
    y = 2
    xy := 1.0
    yx := 2

    fmt.Printf("x=%v, type of %T: ", x,x)
    fmt.Printf("y=%v, type of %T: ", y,y)
    
    fmt.Printf("x=%v, type of %T: ", xy,xy)
    fmt.Printf("y=%v, type of %T: ", yx,yx)

    mean := (x + y )/2.0
    fmt.Printf("Result: %v, type of %T: ", mean, mean)
    fmt.Print(x + y )*/
    
    /*
    Conditional statements

    x:=10

    if x < 5 {
        fmt.Printf("x is small")
    } else{
        fmt.Printf("x is larger")
    }*/
    
    /*
    a :=11.0
    b := 20.0

    if frac := a / b; frac > 0.5{
        fmt.Println("a is more than half of b")
    }*/

    //SWITCH
    n :=2
    switch n {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    case 4: 
        fmt.Println("four")
    default:
        fmt.Println("unknown")
    }
    //Another kind of switch
    switch{
    case n > 100:
        fmt.Println("n is quite big")
    case n > 10:
        fmt.Println("n is big")
    default:
        fmt.Println("n is small")
    }
}