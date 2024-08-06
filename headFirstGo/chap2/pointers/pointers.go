package main

import "fmt"

func main(){
	amount := 6
	double(&amount) // Pass a pointer instead of the variable value
	fmt.Println(amount)
}

func double(number *int){ // Accept a pointer instead of an int value
	*number *= 2
}
