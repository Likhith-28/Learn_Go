// Validate age range 0 - 120 return error otherwise.
package main
import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 0 || age > 120 {
		return errors.New("Invalid age:")
	}
	return nil
}

func main() {
	age := 15
	err := validateAge(age)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Valid age")
}