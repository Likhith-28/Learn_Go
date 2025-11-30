// Copy one Slice to another Slice using Built in Fucntion
package main
import "fmt"
func main(){
	var nums[] int
	nums=append(nums,1,2,3,4,5)
	fmt.Println("Original Slice:",nums)

	nums2:=make([]int,len(nums))
	copy(nums2,nums)
	fmt.Println("Copied Slice:",nums2)
}
