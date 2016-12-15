package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := float64(1); i < 10; i++ {
		fmt.Println("turn %d", i)
		newZ := z - ((z*z) - x)/(2*z)
		if math.Abs(newZ - z) < 0.00000001 {
			fmt.Println("Exit early")
			return newZ
		} else {
			fmt.Println("newZ: %f", newZ)
			z = newZ
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Now real result")
	fmt.Println(math.Sqrt(2.0))
}
