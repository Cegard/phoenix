package client

import (
    "phoenix/messages"
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestClientCreation (t *testing.T) {
    var clientId = 0
    var client = NewClient(clientId)
    
    assert.Equal(t, client.Id, clientId, "Client is not being created")
}


func TestMakeRequest (t *testing.T) {
    var client = NewClient(0)
    var request = messages.NewRequest(client.SetResponse)
    
    assert.Equal(
        t,
        fmt.Sprintf("%T", client.MakeRequest()),
        fmt.Sprintf("%T", request),
        "Client's not creating a request",
    )
}


func TestSetResponse (t *testing.T) {
    var client = NewClient(0)
    var response = messages.NewResponse(0, 0)
    
    client.SetResponse(response)
    
    assert.Equal(
        t,
        fmt.Sprintf("%p", client.GetResponses()[0]),
        fmt.Sprintf("%p", response),
        "Client's not accepting responses",
    )
}


func TestGetREsponses (t *testing.T) {
    var client = NewClient(0)
    var responsesNumber = 10
    
    for i := 0; i < responsesNumber; i++ {
        client.SetResponse(messages.NewResponse(uint(i), 0))
    }
    
    assert.Equal(
        t,
        len(client.GetResponses()),
        responsesNumber,
        "Client is not retrieving all responses",
    )
}
