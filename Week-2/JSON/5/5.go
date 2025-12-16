// Convert JSON string into map[string] interface{}
package main
import (
	"encoding/json"
	"fmt"
)
func main(){
	jsonData := `{"name":"Likhith", "skills":["Go", "Python", "Java"]}`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}