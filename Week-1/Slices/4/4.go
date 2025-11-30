// Carete a Slice from an exsiting Array and modify the values.
package main
import "fmt"
func main(){
	arr:=[5]int{10,20,30,40,50}
	fmt.Println("Array:",arr)
	nums:=arr[:]
	fmt.Println("Slice:",nums)

	nums[0]=100
	nums[1]=200
	fmt.Println("Modified Slice:",nums)
	fmt.Println("Modified Array:",arr)
}