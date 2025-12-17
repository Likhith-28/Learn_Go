// Handle error when opening non-existent file.
package main
import (
	"fmt"
	"os"
)

func main(){
	fileName := "file.txt"
	_, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(fileName)
}