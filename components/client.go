package components

import (
    "sync"
)


type Client struct {
    sync.Mutex
    Id int
    ServerResponses []*Response
}


func (client *Client) MakeRequest () *request {
    
    return NewRequest(client)
}


func (client *Client) SetResponse (response *Response) {
    client.Lock()
    client.ServerResponses = append(client.ServerResponses, response)
    client.Unlock()
}


func CreateClient (id int) *Client {
    
    return &Client{
        Id: id,
        ServerResponses: make([]*Response, 0),
    }
}