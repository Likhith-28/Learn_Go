package main
import "fmt"
func main(){
	var nums []int
	nums=append(nums,1)
	nums=append(nums,2)
	nums=append(nums,3)
	nums=append(nums,4)
	nums=append(nums,5)
	fmt.Println(nums)

	var nums2 []int
	nums2=append(nums2,6)
	nums2=append(nums2,7)
	nums2=append(nums2,8)
	nums2=append(nums2,9)
	nums2=append(nums2,10)
	fmt.Println(nums2)

	nums=append(nums,nums2...)
	fmt.Println(nums)
}