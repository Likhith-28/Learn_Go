// Check if file exists , if not Create it
package main
import(
	"fmt"
	"os"
)

func main(){
	file ,err :=os.Stat("hello.txt")
	if err != nil {
		if os.IsNotExist(err) {
			os.Create("hello.txt")
			return
		}
		fmt.Println("Error :", err)
	}
	fmt.Println("File Created :",file)
}
