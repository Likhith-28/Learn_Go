// Create a Rectangle sruct and computer area
package main
import "fmt"
type Rectangle struct {
	width  float64
	height float64
}
func new(width, height float64) *Rectangle {
	return &Rectangle{
		width: width, 
		height: height,
	}
}
func main(){
	rect := new(10.5, 5.5)
	area := rect.width * rect.height
	fmt.Printf("Area of Rectangle: %.2f\n", area)
}