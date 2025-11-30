package main
import "fmt"
const pi=3.14
func main() {
	// Area of Circle
	var radius float64;
	fmt.Scanf("%f",&radius)
	area:=pi*radius*radius
	fmt.Println("Area of Circle:",area)
}
