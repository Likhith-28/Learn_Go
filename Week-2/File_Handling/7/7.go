// copy content from one file to another
package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	src, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer src.Close()

	dst, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}

	fmt.Println("File copied successfully.")
}