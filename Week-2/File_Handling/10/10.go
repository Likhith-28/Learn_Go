// Write struct data into file in a plain text
package main
import (
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age int
}

func new(Name string, Age int) Person{
	return Person{
		Name : Name,
		Age : Age,
	}
}

func main () {
	person := new("Likhith", 19)
	content := fmt.Sprintf("Name :%s, Age : %d,",person.Name,person.Age)
	file, err := os.Create("Hello.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_,err = file.WriteString(content) 
	if err != nil {
		fmt.Println("Error :",err)
		return
	}
	fmt.Println("File Created Neatly")
}