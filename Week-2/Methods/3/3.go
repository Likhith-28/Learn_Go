// Create Circle struct with Radiys and add a Circumference() method to it.
package main
import "fmt"
type Cricle struct{
	radius float64
}
func new(radius float64) Cricle{
	return Cricle{radius :radius,}
}
func (c Cricle)Circumference() float64{
	return 2 * 3.14 * c.radius
}
func main(){
	var r float64
	fmt.Println("Enter the radius of Cricle:")
	fmt.Scanf("%f", &r)
	circ := new(r)
	fmt.Println("Circumference of Circle:", circ.Circumference())
}