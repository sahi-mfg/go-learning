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

func getFloat() (float64, error){
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil{
		log.Fatal(err)
	}

	var status string
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
