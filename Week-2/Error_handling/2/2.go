// Division function returning error if divsior is zero.
package main
import (
	"errors"
	"fmt"
)
func divide(a int,b int)(int, error){
	if b == 0 {
		return 0, errors.New("Division by zero error")
	}
	return a / b, nil
}
func main(){
	num1 := 10
	num2 := 0
	result, err := divide(num1, num2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
