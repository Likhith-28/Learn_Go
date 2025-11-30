// Count the frequency of elements in a slice using a map.
package main
import "fmt"
func main(){
	dict:=make(map[int]int)
	var nums[] int
	nums=append(nums,1,2,2,3,3,3,4,5,6,6,6,6)
	for _,v:=range nums{
		dict[v]+=1
	}
    fmt.Println("Frequency of elements:", dict)
}