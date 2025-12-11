// Temprature struct with ConvertToFrahenheit method
package main
import "fmt"

type Temprature struct {
	Celsius float64
}
func new (Celsius float64) Temprature {
	return Temprature{Celsius: Celsius,}
}
func (t Temprature) ConvertToFahrenheit() float64 {
	return (t.Celsius * 9 / 5) + 32
}
func main(){
	temp := new(25.45)
	fahrenheit := temp.ConvertToFahrenheit()
	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", temp.Celsius, fahrenheit)
}
