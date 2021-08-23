package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt struct {
	number float64
}

func Sqrt(x float64) (float64, error) {
	if x == 0 {
		return 0, nil
	}
	if x < 0 {

		return 0, ErrNegativeSqrt{x}
	}
	z := 1.0
	for z0 := 0.0; math.Abs(z-z0) > 1e-8; {
		z0 = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.number)
}

func main() {
	fmt.Println(Sqrt(0))
	fmt.Println(Sqrt(1))
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
