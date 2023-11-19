package main

import "fmt"

func main(){
	var a, b int

	fmt.Scan(&a, &b)
	var result = 1

	for i := a; i <= b; i++ {
		result = result * i
	}
	fmt.Println(result)
}