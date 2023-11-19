package main

import "fmt"

func main(){
	var a, b int

	fmt.Scan(&a, &b)
	count := (b-a) +1

	for i := a; i <=b; i++{
		fmt.Println(i)
	} 
	fmt.Println(count)
}