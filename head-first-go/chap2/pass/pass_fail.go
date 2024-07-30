// pass_fail reports whether a grade is passing or failing.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	if input == "100" {
		fmt.Println("Perfect!")
	} else if input >= "60" {
		fmt.Println("You passed!")
	} else if input > "100" {
		fmt.Println("You are a cheater!")
	} else {
		fmt.Println("You failed!")
	}
}
