// Add Perimeter() method to Rectangle struct
package main
import "fmt"

type Rectangle struct{
	length,breadth float64
}
func new(length,breadth float64) Rectangle{
	return Rectangle{length : length,
		breadth : breadth,
	}
}
func (r Rectangle) Perimeter() float64{
	return 2 * (r.length + r.breadth)
}
func main(){
	var l int
	fmt.Println("Enter length of Rectangle:")
	fmt.Scanf("%d", &l)
	var b int
	fmt.Println("Enter breadth of Rectangle:")
	fmt.Scanf("%d", &b)
	rec:=new(float64(l),float64(b))
	fmt.Println("Perimeter of Rectangle:", rec.Perimeter())
}