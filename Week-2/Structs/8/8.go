// Filter students scoring aboe 90 marks
package main
import "fmt"
type Student struct{
	Name string
	Marks int
}
func new(Name string,Marks int)*Student{
	return &Student{
		Name:Name,
		Marks:Marks,
	}
}
func main(){
	s1:=new("Alice",95)
	s2:=new("Bob",85)
	s3:=new("Charlie",92)
	students:=[]*Student{s1,s2,s3}
	var topStudents[]*Student
	for _,student:=range students{
		if student.Marks>90{
			topStudents=append(topStudents,student)
		}
	}
	fmt.Println("Students scoring above 90 marks:")
	for _,student:=range topStudents{
		fmt.Printf("Name: %s, Marks: %d\n",student.Name,student.Marks)
	}
}