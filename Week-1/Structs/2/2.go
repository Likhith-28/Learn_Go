// Create a Car Strut and Print one field using dot notation
package main
import "fmt"
type Car struct{
	Model string
	Year int
}
func new(Model string,Year int)*Car{
	return &Car{
		Model:Model,
		Year:Year,
	}
}
func main(){
	car1:=new("Toyota",2020)
	car2:=new("Honda",2021)
	fmt.Println("Car 1 Model:",car1.Model)
	fmt.Println("Car 2 Model:",car2.Model)
}