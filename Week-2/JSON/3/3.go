// Marshal a slice of structs
package main
import(
	"fmt"
	"encoding/json"
)
type Employe struct{
	Name string `json:"name"`
	Age int `json:"age"`
}
func new(name string, age int) Employe{
	return Employe{
		Name: name,
		Age: age,
	}
}
func main(){
	employees := []Employe{
		new("Likhith",19),
		new("John",21),
		new("Doe",25),
	}
	data, err := json.Marshal(employees)
	if err != nil{
		fmt.Println("Error:", err)
	}else{
		fmt.Println(string(data))
	}
}