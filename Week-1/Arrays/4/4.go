// Count how many even and odd numbers are present in array
package main
import("fmt")

func main(){
	nums := [10]int{10,20,30,40,50,60,70,80,90,100}
	odd := 0
	even := 0

	for num := range nums{
		if num % 2 == 0{
			even++
		}else{
			odd++
		}
	}
	fmt.Println("No.of Even numbers:", even)
	fmt.Println("No.of Odd numbers:", odd)
}