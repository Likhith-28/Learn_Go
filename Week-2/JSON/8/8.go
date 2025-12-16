// marhsal struct with Custom JSON tags
package main
import (
	"encoding/json"
	"fmt"
)

type Employee struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Position string `json:"position"`
}

func new(id int,name string,position string) Employee{
	return Employee{
		ID: id,
		Name: name,
		Position: position,
	}
}

func main(){
	emp := new(101,"Likhith","Developer")
	jsonData, err := json.Marshal(emp)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(string(jsonData))
}