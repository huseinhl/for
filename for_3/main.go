package main

import "fmt"

func main(){
	var a, b int

	fmt.Scan(&a, &b)

	count := (a - b) - 1

	for i := a -1; i > b; i--{
		fmt.Println(i)
	}
	fmt.Println(count)
}	