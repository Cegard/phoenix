package ui

import (
    "phoenix/balancer"
    "phoenix/client"
    "phoenix/utils"
    "testing"
    "time"
    "fmt"
)


func TestProcessSendCommand (t *testing.T) {
    var client = client.NewClient(0)
    var requestsCoef = 10
    var requests = requestsCoef * utils.MaxServiceCapacity
    var command = fmt.Sprintf("send %d", requests)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    ProcessUserCommands(client, command)
    
    if balancer.GetLoadBalancer().TotalRunningInstances() < requestsCoef {
        t.Errorf("Load balancer is not processing the requests")
    }
}
