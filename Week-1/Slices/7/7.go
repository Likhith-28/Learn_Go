// Remove duplicates from a slice while keeping the Order
package main
import "fmt"
func main(){
	var nums[] int
	var n int
	fmt.Scanf("%d",&n)
	for i:=0;i<n;i++{
		var val int
		fmt.Scanf("%d",&val)
		nums=append(nums,val)
	}
}