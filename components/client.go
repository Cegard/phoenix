package components

import (
    "fmt"
)


type Client struct {
    Id int
    serverResponse *Response
}


func (client *Client) MakeRequest () *request {
    
    return NewRequest(client)
}


func (client *Client) SetResponse (response *Response) {
    client.serverResponse = response
    fmt.Println(*client.serverResponse)
}
