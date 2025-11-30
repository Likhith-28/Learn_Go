// Check if number is even or odd
package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d",&n)
	if n%2==0{
		fmt.Println("The Number is even")
	} else {
		fmt.Println("The Number is odd")
	}
}