package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("writing data into a file")
	writeFile("Its fun learning go!!!, Yeah")

	fmt.Println()

	fmt.Println("reading data from a file")
	readFile()

}

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("C:\\Go\\reading-and-writing-to-file\\my-file.txt", bytes, 0644)
	fmt.Println("created a file")
}

func readFile() {
	data, _ := ioutil.ReadFile("C:\\Go\\reading-and-writing-to-file\\my-file.txt")
	fmt.Println("file content: ")
	fmt.Println(string(data))
}
