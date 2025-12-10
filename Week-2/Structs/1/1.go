// Create a Student Struct and Print all fields
package main
import "fmt"
type Student struct{
	Name string
	RollNo int
	Age int
	Grade string
}
func new(Name string,RollNo int,Age int,Grade string)*Student{
	return &Student{
		Name:Name,
		RollNo:RollNo,
		Age:Age,
		Grade:Grade,
	}
}
func main(){
	student1:=new("Alice",101,20,"A")
	student2:=new("Bob",102,21,"B")
	fmt.Println("Student 1:",*student1)
	fmt.Println("Student 2:",*student2)
}