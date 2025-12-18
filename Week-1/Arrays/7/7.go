// Search an element in array using linear search

package main
import("fmt")

func main(){
	nums := [5]int{1,2,3,4,5}
	
	flag := 0
	var key int
	fmt.Scanf("%d", &key)
	for _,num := range nums{
		if num == key{
			flag = 1
		}
	}
	if flag == 0{
		fmt.Println("Element not found !")
	}else{
		fmt.Println("Element found !")
	}
}