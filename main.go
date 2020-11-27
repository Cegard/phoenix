package main

import (
    _"fmt"
    "phoenix/components"
    "sync"
)


func main() {
    var wg sync.WaitGroup
    var loadBalancer = components.GetLoadBalancer(&wg)
    var client = &components.Client{Id: 1}
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    loadBalancer.AssignRequest(client.MakeRequest())
    wg.Wait()
}
