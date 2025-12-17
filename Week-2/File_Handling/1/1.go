// Create a file and wrote text into it
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	defer file.Close()

	text := "Hello, welcome to Go file Handling."
	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	
	fmt.Println("File created and text written successfully")
}


