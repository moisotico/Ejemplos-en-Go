package main 

import (
	"fmt"
	"math"
)


func split(sum int)(int, int){
	x := sum * 4 / 9
	y := sum - x
	return x, y
}


func main() {
	
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Printf("Now we will split a number: \n")

	x:=7
	fmt.Printf("We have chosen the number %v\n", x)
	fmt.Println(split(x))

}

