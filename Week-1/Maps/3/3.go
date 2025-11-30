// Check if a key exists in a map and print the result.
package main
import "fmt"
func main(){
	dict:=make(map[string]int)
	dict["apple"]=5
	dict["banana"]=3
	dict["orange"]=7
	fmt.Println("Initial map:", dict)
	key:="banana"
	_,ok:=dict[key]
	if ok{
		fmt.Println("Key is Exists",key)
	}else{
		fmt.Println("Key is not Exists",key)
	}
}