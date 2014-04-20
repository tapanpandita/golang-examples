package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    fmt.Println("Welcome to the playground!")

    fmt.Println("The time is", time.Now())

    fmt.Println("Or access the network:")
    fmt.Println(net.Dial("tcp", "google.com:80"))
}
