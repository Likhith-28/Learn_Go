// Count number of lines in a file
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	file ,err := os.Open("Hello.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count :=0
	for scanner.Scan() {
		count+=1
	}
	fmt.Println("No.of lines in the file :",count)
}
