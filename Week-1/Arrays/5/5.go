// Find the second largest number in array 
package main
import("fmt")

func main(){
	nums := [5]int{1,2,3,4,5}
	max := nums[0]
	sec := nums[1]
	for _,num := range nums{
		if num > max{
			sec = max
			max = num
		}else if sec != max && num > sec{
			sec = num
		}
	}
	fmt.Println("Second Largest :",sec)
}