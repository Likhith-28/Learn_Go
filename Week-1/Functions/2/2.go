// Wrote divide function which takes two float64 numbers as parameters and returns their division result.
package main
import "fmt"
func divide(a float64,b float64) float64{
	return a/b
}
func main(){
	fmt.Printf("%.2f\n",divide(10.0, 2.3))
}