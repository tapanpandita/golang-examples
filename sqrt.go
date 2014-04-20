package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    /*
    Newton's method of calculating roots. This is a terrible implementation
    which doesn't check for divergence
    */
    z0 := 0.0
    z1 := 1.0

    diff := math.Abs(z0 - z1)

    for diff > 0.001 {
        z0 = z1
        z1 = z1 - ((math.Pow(z1, 2) - x) / 2 * z1)
        diff = math.Abs(z0 - z1)
        //fmt.Println(z0, z1, diff)
    }

    return z1
}

func main() {
    fmt.Println(Sqrt(2))
}
