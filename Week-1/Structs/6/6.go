// Create a Nested Struct
package main 
import "fmt"
type Employee struct{
	name string
	age int
	address Address
}
type Address struct{
	city string
	state string
}
func main(){
	emp := Employee{
		name: "John Doe",
		age: 30,
		address: Address{
			city: "New York",
			state: "NY",
		},
	}
	fmt.Println("Employee Name:", emp.name)
	fmt.Println("Employee Age:", emp.age)
	fmt.Println("Employee City:", emp.address.city)
	fmt.Println("Employee State:", emp.address.state)
}