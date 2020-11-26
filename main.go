package main

import (
    _"fmt"
    "phoenix/components"
    "sync"
)


func main() {
    var wg sync.WaitGroup
    var loadBalancer = components.GetLoadBalancer(&wg)
    var client = &components.Client{Id: 2}
    var request = client.MakeRequest()
    loadBalancer.AssignRequest(request)
    loadBalancer.AssignRequest(client.MakeRequest())
    wg.Wait()
}
