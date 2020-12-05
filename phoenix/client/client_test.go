package client

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestClientCreation (t *testing.T) {
    var clientId = 0
    var client = NewClient(clientId)
    
    assert.Equal(t, client.Id, clientId, "Client is not being created")
}
