// Create a map of student names and marks,then update marks by +5
package main
import "fmt"
func main(){
	dict:=make(map[string]int)
	dict["Likhith"]=85
	dict["Lahari"]=80
	dict["Thasleem"]=90
	dict["Himanshu"]=75
	fmt.Println("Before updating marks:",dict)
	for key,value :=range dict{
		dict[key]=value+5
	}
	fmt.Println("After updating marks:",dict)
}