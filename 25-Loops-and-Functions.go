package main
 
import (
    "fmt"
    "math"
)
 
func Sqrt(x float64) float64 {
    next := float64(1)
    prev := float64(0)
    delta := .01
    for math.Abs(next - prev) > delta {
        prev, next = next, next - (next*next - x)/(2*next)
    }
    return next
}
 
func main() {
    fmt.Println(Sqrt(2))
}