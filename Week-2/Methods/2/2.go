// Add Area() method to Rectangle struct
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
func (r Rectangle) Area() float64{
	return r.length * r.breadth
}
func main(){
	var l int
	fmt.Println("Enter length of Rectangle:")
	fmt.Scanf("%d", &l)
	var b int
	fmt.Println("Enter breadth of Rectangle:")
	fmt.Scanf("%d", &b)
	rec:=new(float64(l),float64(b))
	fmt.Println("Area of Rectangle:", rec.Area())
}