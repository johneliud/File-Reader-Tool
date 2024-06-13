package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check for incorrect number of arguments passed
	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments passed. Usage: go run . input.txt output.txt")
		return
	}

	arguments := os.Args[1:]

	// Check if the file extensions associated with input and output isn't .txt
	if filepath.Ext(arguments[0]) != ".txt" || filepath.Ext(arguments[1]) != ".txt" {
		fmt.Println("Incorrect file extension used on file.")
		return
	}

	inputTxt, err := os.Open(arguments[0])
	if err != nil {
		fmt.Println("Failed opening input file.", err)
		return
	}
	defer inputTxt.Close()

	// Variable inputTxtDetails contains the file properties of the input file using the os.Stat() method
	inputTxtDetails, err := os.Stat(arguments[0])
	if err != nil {
		fmt.Println("Failed to obtain input file properties.", err)
		return
	}

	inputTxtSize := inputTxtDetails.Size()
	if inputTxtSize == 0 {
		fmt.Println("Cannot read from an empty input file.")
		return
	}

	outputTxt, err := os.Create(arguments[1])
	if err != nil {
		fmt.Println("Failed to create output file.", err)
		return
	}

	scanner := bufio.NewScanner(inputTxt)
	for scanner.Scan() {
		readLine := scanner.Text()

		_, err := outputTxt.Write([]byte(readLine))
		if err != nil {
			fmt.Println("Cannot write to output file.", err)
			return
		}
	}

	fmt.Println("File modification successful.")
}
