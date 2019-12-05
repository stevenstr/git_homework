package main

import "fmt"

func main() {
	var hell = "Hello "
	var name string
	fmt.Println("please enter your name:")
	fmt.Scanf("%s", &name)
	fmt.Println(hell, name)
	fmt.Println("How is the weather today?")
}
