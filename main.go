package main

import (
    "phoenix/abstractions"
    "fmt"
)


func main() {
    newRequest := abstractions.CreateRequest(1)
    fmt.Printf("%d \n", newRequest.TimeToProcess)
}