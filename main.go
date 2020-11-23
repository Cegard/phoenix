package main

import (
    "fmt"
    "phoenix/components"
)


func main() {
    var loadBalancer = components.GetLoadBalancer()
    fmt.Println(loadBalancer)
}
