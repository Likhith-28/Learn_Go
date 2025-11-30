package main
import "fmt"
func main(){
	var nums[] int
	nums=append(nums,1)
	nums=append(nums,2)
	nums=append(nums,3)
	nums=append(nums,4)
	nums=append(nums,5)
	fmt.Println(nums)

	index:=2
	nums=append(nums[:index],nums[index+1:]...)
	fmt.Println(nums)
}