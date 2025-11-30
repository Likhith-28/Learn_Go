package main
// Convert int to float64, float64 to int.
import "fmt"

func main(){
	var intVal int = 42
	var floatVal float64 = 42.58

	// Convert int to float64
	convertedFloat := float64(intVal)
	fmt.Printf("Converted int to float64: %f\n", convertedFloat)

	// Convert float64 to int
	convertedInt := int(floatVal)
	fmt.Printf("Converted float64 to int: %d\n", convertedInt)
}