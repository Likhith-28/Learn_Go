// Return a new slice with unique values
package main
import "fmt"
func unique(nums []int)[]int{
	seen:=make(map[int]bool)
	var unique[] int
	for _,v := range nums{
		if !seen[v]{
			seen[v]=true
			unique=append(unique,v)
		}
	}
	return unique
}
func main(){
	var nums[] int
	var n int
	fmt.Scanf("%d",&n)
	for i:=0;i<n;i++{
		var val int
		fmt.Scanf("%d",&val)
		nums=append(nums,val)
	}
	fmt.Println("Original Slice:",nums)
	uniqueSlice:=unique(nums)
	fmt.Println("Slice with unique values:",uniqueSlice)
}
