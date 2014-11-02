package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("Sqrt negative number: %g", e)
}

func Sqrt(f float64) (float64, error) {
    next := float64(1)
    prev := float64(0)
    delta := .01
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    for math.Abs(next - prev) > delta {
        prev, next = next, next - (next*next - f)/(2*next)
    }
    return next, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}