// Return error if number is negative.
package main
import (
	"errors"
	"fmt"
)

func check(num int) error {
	if num < 0 {
		return errors.New("negative number error")
	}
	return nil
}

func main(){
	nums := -5
	err := check(nums)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(nums)
}