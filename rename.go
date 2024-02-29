// sources:
// 1) https://www.geeksforgeeks.org/fmt-scanln-function-in-golang-with-examples/
// 2)
// 3)
// 4)
package main

import (
	"fmt"
	"os"
)

func main() {
	var directory string
	var filename string
	var filetype string
	fmt.Println("where is the file?")
	fmt.Scanln(&directory)
	fmt.Println("what is its name?")
	fmt.Scanln(&filename)
	fmt.Println("what is the file's type?")
	fmt.Scanln(&filetype)

	find_file(directory, filename, filetype)
}

func find_file(directory string, filename string, filetype string) {
	_, err := os.Stat(directory + filename + filetype)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			// Handle other errors
			fmt.Println("Error checking file:", err)
		}
	} else {
		// File exists
		fmt.Println("File exists")
	}
}
