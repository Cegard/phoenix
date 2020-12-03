package client

import (
    "phoenix/messages"
    "sync"
)


type Client struct {
    sync.Mutex
    Id int
    ServerResponses []*messages.Response
}


func (client *Client) MakeRequest() *messages.Request {
    
    
    var respondToThis = func (response *messages.Response) {
        client.ServerResponses = append(client.ServerResponses, response)
    }
    
    
    return messages.NewRequest(respondToThis)
}


func (client *Client) SetResponse (response *messages.Response) {
    client.Lock()
    client.ServerResponses = append(client.ServerResponses, response)
    client.Unlock()
}


func NewClient (id int) *Client {
    
    return &Client {
        Id: id,
        ServerResponses: make([]*messages.Response, 0),
    }
}
