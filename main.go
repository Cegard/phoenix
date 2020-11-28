package main

import (
    "phoenix/components"
    "sync"
)


func main() {
    var wg sync.WaitGroup
    var loadBalancer = components.GetLoadBalancer(&wg)
    var client = &components.Client{Id: 1}
    
    for i := 0; i < 500; i++ {
        loadBalancer.AssignRequest(client.MakeRequest())
    }
    
    wg.Wait()
}
