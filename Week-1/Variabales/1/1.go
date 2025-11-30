package main
// Declare variables of all basic types and print their types.
import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 10
	var f float64 = 20.5
	var b bool = true
	var s string = "Hello, Go!"

	fmt.Printf("Type of i: %s, Value: %d\n", reflect.TypeOf(i), i)
	fmt.Printf("Type of f: %s, Value: %f\n", reflect.TypeOf(f), f)
	fmt.Printf("Type of b: %s, Value: %t\n", reflect.TypeOf(b), b)
	fmt.Printf("Type of s: %s, Value: %s\n", reflect.TypeOf(s), s)
}
