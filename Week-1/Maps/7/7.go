// Delete all keys whose value is below the threshold value.
package main
import "fmt"
func main(){
	dict:=make(map[string]int)
	dict["apple"]=5
	dict["banana"]=3
	dict["orange"]=7
	dict["mango"]=10
	dict["grape"]=4
	fmt.Println("Initial map:", dict)

	threshold:=5
	for key,value:=range dict{
		if value<threshold{
			delete(dict,key)
		}
	}
	fmt.Println("Map after deletion:", dict)
}