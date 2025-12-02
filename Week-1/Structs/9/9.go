// Create an Address struct and embed it int a Person struct
package main 
import "fmt"
type Perosn struct{
	name string
	age int
	address Address
}
type Address struct{
	city string
	state string
}
func main(){
	per := Perosn{
		name: "John Doe",
		age: 30,
		address: Address{
			city: "New York",
			state: "NY",
		},
	}
	fmt.Println("Person Name:", per.name)
	fmt.Println("Person Age:", per.age)
	fmt.Println("Person City:", per.address.city)
	fmt.Println("Person State:", per.address.state)
}