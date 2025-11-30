// Itreate over a map and print all key-value pairs.
package main
import "fmt"
func main(){
	dict:=make(map[string]int)
	dict["apple"]=5
	dict["banana"]=3
	dict["orange"]=7
	fmt.Println("Initial map:", dict)
	for key,value := range dict {
		fmt.Println(key,":",value)
	}
}