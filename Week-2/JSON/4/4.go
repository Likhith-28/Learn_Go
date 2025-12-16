// Unmarshal JSON array into slice of structs
package main
import (
	"encoding/json"
	"fmt"
)
type Employe struct{
	Name string `json:"name"`
	Skills []string `json:"skills"`
}
func main(){
	jsonData := `[{"name":"Likhith", "skills":["Go", "Python", "Java"]},
	{"name":"Lahari", "skills":["Go", "Python", "Java"]},
	{"name":"Arjun", "skills":["Go", "Python", "Java"]}]`
	var e []Employe
	err := json.Unmarshal([]byte(jsonData),&e)
	if err != nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println(e)
}