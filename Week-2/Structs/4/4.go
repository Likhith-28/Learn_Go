// Modify an Employee struct's salary using a pointer 
package main
import "fmt"
type Employee struct{
	Name string
	Salary float64
}
func new(Name string,Salary float64) *Employee{
	return &Employee{
		Name:Name,
		Salary:Salary,
	}
}
func (e *Employee) modifysalary(sal float64){
	e.Salary=sal
}
func main(){
	emp:=new("Likhith",4500.95)
	fmt.Println("Before modification Salary:",emp.Salary)
	emp.modifysalary(5000.00)
	fmt.Println("After modification Salary:",emp.Salary)
}