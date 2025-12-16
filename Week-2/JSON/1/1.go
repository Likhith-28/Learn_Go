// Marshal a Person struct to JSON
package main
import (
	"fmt"
	"encoding/json"
)
type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func New (name string, age int) Person{
	return Person{
		Name: name,
		Age: age,
	}
}
func main(){
	p := New("Likhith",19)
	jsonData, err := json.Marshal(p)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}