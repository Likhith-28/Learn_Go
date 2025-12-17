// Search element in slice; return error if not found.
package main
import (
	"errors"
	"fmt"
)

func Search(slice []int, target int) (int, error) {
	for i := range slice {
		if slice[i] == target {
			return i, nil
		}
	}
	return -1,errors.New("element not found")
}

func main() {
	slice := []int{10,20,30,40,50}
	target := 35
	index, err := Search(slice, target)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Element found at index: %d\n", index)
	
}