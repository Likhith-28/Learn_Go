// Read file line by line using bufio.Scanner
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){ // Moves to the next line
		fmt.Println(scanner.Text()) // Prints the current line
	}
}