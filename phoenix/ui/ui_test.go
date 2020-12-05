package ui

import (
    "phoenix/balancer"
    "phoenix/client"
    "phoenix/utils"
    "time"
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestProcessSendCommand (t *testing.T) {
    var client = client.NewClient(0)
    var requestsCoef = 10
    var requests = requestsCoef * utils.MaxServiceCapacity
    var command = fmt.Sprintf("send %d", requests)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    ProcessUserCommands(client, command)
    
    assert.GreaterOrEqual(
        t,
        requestsCoef,
        balancer.GetLoadBalancer().TotalRunningInstances(),
        "Load balancer is not processing the requests",
    )
}
