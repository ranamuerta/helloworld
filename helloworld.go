package main

import(
	"fmt"
	"os"
)

func main(){

	fmt.Println(hw())

	env()
}

func hw() string{
	return "Hello world"
}

func env(){
	for _, env := range os.Environ(){
		fmt.Println("env: ", env)
	}
}