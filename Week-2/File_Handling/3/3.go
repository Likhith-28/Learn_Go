// Append text to an exisiting file
package main
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("hello.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nAppended line of text.")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text appended successfully.")
}
