// Delete a file using os.Remove()
package main
import (
	"fmt"
	"os"
)

func main(){
	err := os.Remove("hello.txt")
	if err != nil {
		fmt.Println("Error :",err)
		return
	}
	fmt.Println("File Deleted")
}
