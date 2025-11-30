// Take user input (name, age) and print a formatted message.
package main
import "fmt"
func main(){
	var name string
	var age int
	fmt.Println("Enter your Name:")
	fmt.Scanf("%s",&name)
	fmt.Println("Enter your Age:")
	fmt.Scanf("%d",&age)
	fmt.Printf("Hello, %s! You are %d years old.\n",name,age)
}