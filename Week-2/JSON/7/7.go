// Unmarshal nested JSON (student + subjects list)
package main
import (
	"encoding/json"
	"fmt"
)

type Student struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Subjects []string `json:"subjects"`
}
func main(){
	jsonData := `[
		{"name": "Alice", "age": 20, "subjects": ["Math", "Science"]},
		{"name": "Bob", "age": 22, "subjects": ["History", "Art"]}
	]`

	var students []Student
	
	err := json.Unmarshal([]byte(jsonData), &students)
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(students)
}
