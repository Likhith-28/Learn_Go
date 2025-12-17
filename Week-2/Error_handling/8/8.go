// Validate email; return error if "@" is missing.
package main
import (
	"errors"
	"fmt"
	"strings"
)

func validateEmail(email string) error {
	if !strings.Contains(email, "@"){
		return errors.New("Invalid email : missing @ symbol")
	}
	return nil
}

func main() {
	email := "pollailikhith.gmail.com"
	err := validateEmail(email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Valid email address")
}
