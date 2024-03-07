// sources:
// 1) https://www.geeksforgeeks.org/fmt-scanln-function-in-golang-with-examples/
// 2) https://pkg.go.dev/os
// 3)
// 4)
package main

import (
	"fmt"
	"os"
)

func main() {
	var running bool = true
	for running != false {

		var userInput string
		fmt.Println("menu: \n 1: rename a single file \n 2: rename all the files in a folder \n 3: exit")
		fmt.Scanln(&userInput)
		switch userInput {
		case "1":
			rename_inputs()
		case "2":
			fmt.Println("2")
		case "3":
			running = false
		}
	}
}

func rename_inputs() {
	// declare variables
	var directory string
	var filename string
	var filetype string
	var newName string

	// get user input
	fmt.Println("where is the file?")
	fmt.Println("enter the file path to the folder containing the file(s)")
	fmt.Scanln(&directory)
	fmt.Println("what is its name?")
	fmt.Scanln(&filename)
	fmt.Println("what is the file's type?")
	fmt.Scanln(&filetype)
	fmt.Println("what is the file's new name? ")
	fmt.Scanln(&newName)

	find_file(directory, filename, filetype)
	rename_single_file(directory, filename, filetype, newName)
}

func read_directory(filepath string) {
	entries, err := os.ReadDir(filepath)
	// if it cant read the folder
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		// Check if it's not a folder
		if !entry.IsDir() {
			fmt.Println(entry.Name()) // Print the file name
		}
	}
}

func rename_single_file(directory, filename, filetype, newName string) {

	// Rename the file
	err := os.Rename(directory+"/"+filename+filetype, directory+"/"+newName+filetype)
	if err != nil {
		fmt.Println("Error renaming file:", err)
	}

	// Check if the file was renamed successfully
	if _, err := os.Stat(newName); os.IsNotExist(err) {
		fmt.Println("File was not renamed successfully")
	} else {
		fmt.Println("File was renamed successfully")
	}
}

func find_file(directory string, filename string, filetype string) {
	//_ dennotes that the var is not used if you need to use it replace it with filename
	_, err := os.Stat(directory + "/" + filename + filetype)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			// Handle other errors
			fmt.Println("Error checking file:", err)
		}

	} else {

		// File exists
		// return fileInfo
		fmt.Println("File exists")
		// fmt.Println("file name: ", fileInfo.Name())
		// fmt.Println("File size:", fileInfo.Size())
		// fmt.Println("Is directory:", fileInfo.IsDir())
		// fmt.Println("Last modified:", fileInfo.ModTime())
	}
}
