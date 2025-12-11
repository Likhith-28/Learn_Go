// Student struct with Grade() method returing A/B/C
package main
import "fmt"
type Student struct{
	name string
	marks int
}
func new(name string,marks int)Student{
	return Student{name : name,
	marks : marks,}
}
func Grade(s Student)string{
	if s.marks>=90{
		return "A"
	}else if s.marks>=75{
		return "B"
	}else{
		return "C"
	}
}

func main(){
	s:=new("Likhith",85)
	fmt.Println("Student Name:",s.name)
	fmt.Println("Student Marks:",s.marks)
	fmt.Println("Student Grade:",Grade(s))
}