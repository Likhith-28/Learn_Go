// Create a struct with JOSN tags and print values 
package main
import (
	"encoding/json"
	"fmt"
)
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}
func main(){
	jsonData := `{"name":"Jane Doe", "age":25, "address":"123 Main St"}`
	var person Person
	err := json.Unmarshal([]byte(jsonData),&person)
	if err!=nil{
		fmt.Println("Error Loading JSON:",err)
		return
	}
	fmt.Println("Person Name:", person.Name)
	fmt.Println("Person Age:", person.Age)
	fmt.Println("Person Address:", person.Address)
}