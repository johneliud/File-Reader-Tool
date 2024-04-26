package main

import (
	"fmt"
	"os"
)

func main() {
	// Check for incorrect number of arguments passed
	if len(os.Args) != 3 {
		fmt.Println("Error! Incorrect number of arguments passed. Usage: go run . input-file.txt output-file.txt")
		return
	}

	arguments := os.Args[1:]

	// inputFile stores the contents of the read file 
	inputFile, err := os.ReadFile(arguments[0])
	if err != nil {
		fmt.Println("Error! Cannot read from file.")
		return
	}

	// Writes whatever is read from inputFile and specifies file permissions of the new file
	err = os.WriteFile(arguments[1], []byte(inputFile), 0o644)
	if err != nil {
		fmt.Println("Error! Cannot write to file")
		return
	}

	fmt.Println("File modification successful")

}
