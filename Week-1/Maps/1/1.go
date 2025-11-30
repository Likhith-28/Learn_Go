// Create a Map of fruits and prices,then add and delete some keys
package main
import "fmt"
func main(){
	fruits:=make(map[string]float64)
	fruits["apple"]=1.2
	fruits["banana"]=0.5
	fruits["orange"]=0.8
	fmt.Println("Fruits map:", fruits)

	// Adding a new key-value pair
	fruits["grape"]=2.0
	fmt.Println("After adding grape:", fruits)

	// Deleting a key-value pair
	delete(fruits, "banana")
	fmt.Println("After deleting banana:", fruits)
}