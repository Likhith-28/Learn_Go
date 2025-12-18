// Rotate an array to the right by 1 position

package main
import("fmt")

func main(){
	nums := [5]int{10,20,30,40,50}
	fmt.Println("Before Rotation :",nums)
	temp := nums[len(nums)-1]
	for i := len(nums)-2; i > 0; i--{
		nums[i+1] = nums[i]
	}
	nums[0] = temp
	fmt.Println("After Rotation by 1 Position:",nums)
}