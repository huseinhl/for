package main

import "fmt"

func main(){
	var price float64 

	fmt.Scan(&price)

	for i := 0.1; i <= 1; i+=0.1{
		fmt.Printf("%.2f\n", float64(i)*price)
	}
}