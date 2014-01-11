package main

import (
	"fmt"
	"github.com/cstrail/newmath"
	//"math"
)

func main() {

	for i := 1; i <= 100; i++ {
		b := float64(i)
		fmt.Printf("Sqrt(%d) = %f\n", i, newmath.Sqrt(b))

		if i%10 == 0 {
			fmt.Println()
		}

	}
	//fmt.Printf("Hello, world. Sqrt(5) = %v\n", newmath.Sqrt(5))
	fmt.Printf("\nThis is the end of the Square Root list.\n\n")
}
