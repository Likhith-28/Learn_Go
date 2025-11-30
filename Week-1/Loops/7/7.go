// Print the Multiplication Table of bu taking user input
package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter the Table Number:")
	fmt.Scanf("%d",&n)
	for i:=1;i<=10;i++{
		fmt.Println(n," x ",i," = ",n*i)
	}
}