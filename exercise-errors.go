package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for i := float64(1); i < 10; i++ {
		//fmt.Println("turn %d", i)
		newZ := z - ((z*z) - x)/(2*z)
		if math.Abs(newZ - z) < 0.00000001 {
			//fmt.Println("Exit early")
			return newZ, nil
		} else {
			//fmt.Println("newZ: %f", newZ)
			z = newZ
		}
	}
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g",
		float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}