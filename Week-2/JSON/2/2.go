// UnMarshal Json string to struct
package main
import (
	"fmt"
	"encoding/json"
)
type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func main(){
	jsonString := `{"name":"Likhith","age":19}`
	var p Person
	err := json.Unmarshal([]byte(jsonString),&p)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(p)
}