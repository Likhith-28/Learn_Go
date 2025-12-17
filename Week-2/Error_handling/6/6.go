// wrap underlying error using fmt.Errorf with %w
package main
import (
	"errors"
	"fmt"
)

func readFile(filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}
	return fmt.Errorf("failed to read file %s: %w", filename, errors.New("file not found"))
}

func main() {
	err := readFile("Hello.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println("Example completed successfully.")
}