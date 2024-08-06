package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("Enter a file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	fileInfo, error := os.Stat(fileName)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is directory:", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
}
