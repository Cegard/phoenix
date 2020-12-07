package client

import (
    "phoenix/messages"
    "sync"
)


type Client struct {
    sync.Mutex
    Id int
    serverResponses []*messages.Response
}


func NewClient (id int) *Client {
    
    return &Client {
        Id: id,
        serverResponses: make([]*messages.Response, 0),
    }
}


func (client *Client) MakeRequest() *messages.Request {
    
    return messages.NewRequest(client.SetResponse)
}


func (client *Client) SetResponse (response *messages.Response) {
    client.Lock()
    client.serverResponses = append(client.serverResponses, response)
    client.Unlock()
}


func (client *Client) GetResponses() []*messages.Response {
    client.Lock()
    defer client.Unlock()
    
    return client.serverResponses
}
