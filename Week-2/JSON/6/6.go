// Convert struct JSON and back to struct
package main
import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skills []string `json:"skills"`
}

func new(name string, age int, skills []string) Person{
	return Person{
		Name: name,
		Age: age,
		Skills: skills,
	}
}

func main(){
	p1 := new("Likhith", 19, []string{"Go", "Python", "JavaScript"})

	//Struct to JSON
	jsonData, err := json.Marshal(p1)
	if err != nil{
		fmt.Println("Error :",err)
		return
	}
	fmt.Println("Json Data :",string(jsonData))

	//JSON to Struct
	var p2 Person
	err = json.Unmarshal([]byte(jsonData), &p2)
	if err != nil{
		fmt.Println("Error :",err)
		return
	}
	fmt.Println("Struct Data :",p2)
}
