// Check if array is sorted or not

package main
import("fmt")

func main(){
	nums := [5]int{10,20,80,40,50}
	flag := false
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] < nums[i-1]{
			flag = true
		}
	}

	if flag{
		fmt.Println("Array is not Sorted")
	}else{
		fmt.Println("Array is Sorted")
	}
}
