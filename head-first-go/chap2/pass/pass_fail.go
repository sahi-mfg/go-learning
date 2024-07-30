// pass_fail reports whether a grade is passing or failing.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var status string
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	input = strings.TrimSpace(input)
	grade, error := strconv.ParseFloat(input, 64)
	if error != nil {
		log.Fatal(error)
	}

	if grade > float64(100) {
		status = "Not Possible"
	} else if grade >= float64(60) {
		status = "Passing"
	} else if grade == float64(100) {
		status = "Passing with distinction"
	} else {
		status = "Failing"
	}
	fmt.Println("A grade of", grade, "out of 100", "is", status)
}
