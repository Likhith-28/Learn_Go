// Handle optional/missing fields in JSON unmarshalling

package main
import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city,omitempty"`
}

func main(){
	jsonData := `{"name":"Likhith","age":19}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Name: %s\n", person.Name)
	fmt.Printf("Age: %d\n", person.Age)
	fmt.Printf("City: %s\n", person.City) // City will be empty if not present in JSON
}
