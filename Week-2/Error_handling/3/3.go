// Convert string to int using strconv.Atoi and handle error.
package main
import (
	"fmt"
	"strconv"
)

func main(){
	strNum := "123ab"
	num, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Number:", num)
}