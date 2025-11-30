// Find the Most expensive fruit from a fruit-price map
package main
import "fmt"
func main(){
	dict:=make(map[string]int)
	dict["apple"]=5
	dict["banana"]=3
	dict["orange"]=7
	dict["mango"]=10
	dict["grape"]=4

	max:=0
	for _,value :=range dict{
		if value>max{
			max=value
		}
	}
	fmt.Println("Most expensive fruit price is:",max)
}