// Check if a string is a palindrome using a function.
package main
import "fmt"
func check(s string)bool{
	n:=len(s)
	for i:=0;i<n/2;i++{
		if s[i]!=s[n-i-1]{
			return false
		}
	}
	return true
}
func main(){
	fmt.Println(check("madam"))
	fmt.Println(check("hello"))
}