// read a file and print its content
package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	text := "Hello, welcome to Go file Handling."
	text2 := "\nThis is the second line added to the file."
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	_, err = file.WriteString(text2)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	
	fmt.Println("File created and text written successfully")

	data, err := os.ReadFile("Hello.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	
	fmt.Println("File content:")
	fmt.Println(string(data))
}

