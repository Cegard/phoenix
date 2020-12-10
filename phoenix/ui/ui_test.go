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


func assertEmptyAndNil (t *testing.T, message string, err error) {
    assert.Empty(
        t,
        message,
        "Shouldn't be a message here",
    )
    assert.Nil(
        t,
        err,
        "Should be nil",
    )
}


func assertNotEmptyAndNil (t *testing.T, message string, err error) {
    assert.NotEmpty(
        t,
        message,
        "Should be a message here",
    )
    assert.Nil(
        t,
        err,
        "Should be nil",
    )
}


func assertEqualAndError (t *testing.T, message string, msgTocompare string, err error) {
    assert.Equal(
        t,
        message,
        msgTocompare,
        "Messages don't correspond",
    )
    assert.Error(
        t,
        err,
        "Should return an error here",
    )
}


func testSendCommand (badCommand string, client *client.Client, t *testing.T) {
    var requestsCoef = 10
    var requests = requestsCoef * utils.MaxServiceCapacity
    var goodCommand = fmt.Sprintf(badCommand, requests)
    
    message, err := ProcessUserCommands(client, badCommand)
    assertEqualAndError(t, message, utils.CouldNotProcessMsg, err)
    
    message, err = ProcessUserCommands(client, goodCommand)
    
    assertEmptyAndNil(t, message, err)
    assert.GreaterOrEqual(
        t,
        requestsCoef,
        balancer.GetLoadBalancer().TotalRunningInstances(),
        "Load balancer is not processing the requests",
    )
}


func testServiceStatusCommand (badCommand string, client *client.Client, t *testing.T) {
    var goodCommand = fmt.Sprintf(badCommand, 1)
    
    message, err := ProcessUserCommands(client, goodCommand)
    assertNotEmptyAndNil(t, message, err)
    
    message, err = ProcessUserCommands(client, fmt.Sprintf(badCommand, 0))
    assertEmptyAndNil(t, message, err)
    
    message, err = ProcessUserCommands(client, badCommand)
    assertEqualAndError(t, message, utils.CouldNotProcessMsg, err)
}


func testBalancerStatusCommand (client *client.Client, t *testing.T) {
    message, err := ProcessUserCommands(client, "serverStatus")
    assertNotEmptyAndNil(t, message, err)
}


func testDefaultAnswer (client *client.Client, t *testing.T) {
    message, err := ProcessUserCommands(client, "")
    
    assert.Equal(
        t,
        message,
        utils.NotRecognizedMsg,
        "Messages don't correspond",
    )
    assert.Nil(
        t,
        err,
        "Shouldn't be an error on processing a non existing command",
    )
}


func TestProcessUserCommands (t *testing.T) {
    var client = client.NewClient(0)
    var sendCommand = fmt.Sprintf("%s", utils.Send)
    var serviceStatusCommand = fmt.Sprintf("%s", utils.ServiceStatus)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    testSendCommand(sendCommand + " %d", client, t)
    testServiceStatusCommand(serviceStatusCommand + " %d", client, t)
    testBalancerStatusCommand(client, t)
    testDefaultAnswer(client, t)
}
