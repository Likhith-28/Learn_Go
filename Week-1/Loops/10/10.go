package main
import "fmt"
func main(){
	var n int
	fmt.Scanf("%d",&n)
	temp:=n
	sum:=0
	for temp>0{
		digit:=temp%10
		sum=(sum*10)+digit
		temp/=10
	}
	fmt.Println("The Reverse of digits:",sum)
}