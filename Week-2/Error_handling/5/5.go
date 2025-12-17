// Create custom errr using errors.New()
package main
import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("This is a custom error message")
	if err != nil {
		fmt.Println("Error:", err)
	}
}