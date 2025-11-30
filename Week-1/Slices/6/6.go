// Find the length and capticy of a lsice after multiple appends
package main
import "fmt"
func main(){
	var nums []int
	fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
	for i:=0; i<10; i++ {
		nums = append(nums, i)
		fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
	}
}

