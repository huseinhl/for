package main

import "fmt"

func main(){
	var n float64

	fmt.Scan(&n)

	for i := 1.0; i <= n; i+=0.1 {
		fmt.Printf("%.1f\n", float64(i)*n)
	}
	
}