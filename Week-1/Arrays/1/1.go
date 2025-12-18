// Find the maximum, minimum, and sum of elements in an array.
package main
import ("fmt")

func main(){
	nums := [5]int {10,20,30,40,50}
	var sum int
	min := nums[0]
	max := nums[0]
	for _, num := range nums{
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	fmt.Println("Sum :",sum)
	fmt.Println("Min :",min)
	fmt.Println("Max :",max)
}
