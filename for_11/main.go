package main

import "fmt"

func main(){
	var n int

	fmt.Scan(&n)
	var result int 

	for i := n; i <= (2*n)*(2*n); i +=1 {
		fmt.Println(i)
		result += (n+i)*(n+i)
	}
	fmt.Println(result)
}