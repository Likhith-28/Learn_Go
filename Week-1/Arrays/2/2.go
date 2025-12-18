// Reverse an array without using an extra array
package main
import ("fmt")
func main(){
	nums := [5]int {10,20,30,40,50}
	var temp int
	i := 0
	j := len(nums) - 1
	
	for i < j{
		temp = nums[i]
		nums[i] = nums[j]
		nums[j] = temp

		i++
		j--
	}

	fmt.Println(nums)
}