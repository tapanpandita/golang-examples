package main

import "fmt"

func main() {
    for i := 0; i < 25; i++ {

        if i % 3 == 0 {
            fmt.Printf("Fizz")
        }

        if i % 5 == 0 {
            fmt.Printf("Buzz")
        }

        if (i % 5 == 0) || (i % 3 == 0) {
            fmt.Printf("\n")
        }
    }
}
