package main

import (
	"fmt"
	"math/cmplx"
)

const delta = 1e-10

func Cbrt(x complex128) complex128 {
	z := x
	for {
		n := z - (cmplx.Pow(z, 3)-x)/(3* cmplx.Pow(z, 2))
		if cmplx.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z
}

func main() {
	const x = 2
	mine, theirs := Cbrt(x), cmplx.Pow(x, 1./3.)
	fmt.Println(mine, theirs, mine-theirs)
}