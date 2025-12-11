// Counter struct with Increment() And Decrement() methods
package main
import "fmt"
type Counter struct{
	value int
}
func new(val int) Counter{
	return Counter{value:val}
}
func (c *Counter)Increment(val int){
	c.value+=val
}
func (c *Counter)Decrement(val int){
	c.value-=val
}
func main(){
	c:=new(10)
	fmt.Println("Initial Value:",c.value)
	c.Increment(5)
	fmt.Println("After Incrementing by 5:",c.value)
	c.Decrement(3)
	fmt.Println("After Decrementing by 3:",c.value)
}