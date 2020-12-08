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
    var asserter = assert.New(t)
    var client = client.NewClient(0)
    var requestsCoef = 10
    var requests = requestsCoef * utils.MaxServiceCapacity
    var badCommand string
    var goodCommand string
    var message string
    var err error
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    // Testing 'send #' command
    badCommand = "send %d"
    goodCommand = fmt.Sprintf(badCommand, requests)
    message, err = ProcessUserCommands(client, goodCommand)
    
    asserter.GreaterOrEqual(
        requestsCoef,
        balancer.GetLoadBalancer().TotalRunningInstances(),
        "Load balancer is not processing the requests",
    )
    asserter.Empty(
        message,
        "Wrong message",
    )
    asserter.Nil(
        err,
        "Should be nil",
    )
    
    message, err = ProcessUserCommands(client, badCommand)
    
    asserter.Equal(
        message,
        utils.CouldNotProcessMsg,
        "Messages don't correspond",
    )
    asserter.Error(
        err,
        "Not an error",
    )
    
    // Testing 'serverStatus' command
    goodCommand = "serverStatus"
    message, err = ProcessUserCommands(client, goodCommand)
    
    asserter.NotEmpty(
        message,
        "Wrong message",
    )
    asserter.Nil(
        err,
        "Should be nil",
    )
    
    // Testing 'serviceStatus #' command
    badCommand = "serviceStatus %d"
    goodCommand = fmt.Sprintf(badCommand, 1)
    message, err = ProcessUserCommands(client, goodCommand)
    
    asserter.NotEmpty(
        message,
        "Wrong message",
    )
    asserter.Nil(
        err,
        "Should be nil",
    )
    
    message, err = ProcessUserCommands(client, badCommand)
    
    asserter.Equal(
        message,
        utils.CouldNotProcessMsg,
        "Messages don't correspond",
    )
    asserter.Error(
        err,
        "Not an error",
    )
    
    // Testing default answer
    message, err = ProcessUserCommands(client, "")
    
    asserter.Equal(
        message,
        utils.NotRecognizedMsg,
        "Messages don't correspond",
    )
    asserter.Nil(
        err,
        "Not an error",
    )
}
