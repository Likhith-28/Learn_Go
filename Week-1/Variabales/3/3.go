package main
import "fmt"
func main(){
	var a,b int
	fmt.Println("Enter Number 1:")
	fmt.Scanf("%d",&a)
	fmt.Println("Enter Number 2:")
	fmt.Scanf("%d",&b)
	fmt.Println("Before Swapping:", a, b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Println("After Swapping",a,b)
}