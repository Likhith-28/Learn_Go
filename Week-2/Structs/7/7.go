// Create a slice of student structs and print all
package main
import "fmt"
type Student struct{
	Name string
	Rollno int
	Age int
	Grade string
}
func new(Name string,Rollno int,Age int,Grade string) *Student{
	return &Student{
		Name:Name,
		Rollno:Rollno,
		Age:Age,
		Grade:Grade,
	}
}
func main(){
	s1:=new("Alice",1,20,"A")
	s2:=new("Bob",2,21,"B")
	s3:=new("Charlie",3,22,"A")
	Students:=[] Student{}
	Students=append(Students,*s1)
	Students=append(Students,*s2)
	Students=append(Students,*s3)
	for _,student:=range Students{
		fmt.Println(student)
	}
}
