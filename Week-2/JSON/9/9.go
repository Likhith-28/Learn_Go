// Pretty-Print JSON with identation
package main
import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func new(name string, age int, city string) Person{
	return Person{
		Name: name,
		Age:  age,
		City: city,
	}
}

func main(){
	person := new("Likhith",19,"Nellore")
	jsonData, err := json.MarshalIndent(person,""," ")
	if err != nil {
		fmt.Println("Error :",err)
		return
	}

	fmt.Println(string(jsonData))
}