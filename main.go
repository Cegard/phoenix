package main

import (
    _"fmt"
    "phoenix/components"
)


func main() {
    var loadBalancer = components.GetLoadBalancer()
    var client = &components.Client{Id: 2}
    var request = client.MakeRequest()
    loadBalancer.AssignRequest(request)
}
