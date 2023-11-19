package main

import "fmt"

func main(){
	var a, b int

	fmt.Scan(&a, &b)
	var result int

	for i:=a; i <=b; i++{
		fmt.Println(i)
		result += i
	}
	fmt.Println(result)

}