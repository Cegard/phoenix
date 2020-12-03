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
    var requests = requestsCoef * utils.MAX_SERVICE_CAPACITY
    var command = fmt.Sprintf("send %d", requests)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MAX_PROCESS_TIME))
    
    ProcessUserCommands(client, command)
    
    if len(balancer.GetLoadBalancer().services) < requestsCoef {
        t.Errorf("Load balancer is not processing the requests")
    }
}
