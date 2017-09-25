package main

import(
	"fmt"
	"os"
)

func main(){

	fmt.Println("Hello World")

	for _, env := range os.Environ(){
		fmt.Println("env: ", env)
	}
}
