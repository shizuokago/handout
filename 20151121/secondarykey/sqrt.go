package main

import (
	"math"
)

func Sqrt1(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2.0)-x)/(2*z)
	}
	return z
}

func Sqrt2(x float64) float64 {
	z := x
	for {
		prez := z
		z = z - (math.Pow(z, 2.0)-x)/(2*z)
		if math.abs(prez-z) < 1e-10 {
			break
		}
	}
}
