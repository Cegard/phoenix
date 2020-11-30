package components

import (
    "testing"
)


func TestClientCreation (t *testing.T) {
    var clientId = 0
    var client = CreateClient(clientId)
    
    if client.Id != clientId {
        t.Errorf("Client is not being created")
    }
}
