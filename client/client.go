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
    
    return messages.NewRequest(client.SetResponse)
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
