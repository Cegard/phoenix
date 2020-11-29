package components


type Client struct {
    Id int
    ServerResponses []*Response
}


func (client *Client) MakeRequest () *request {
    
    return NewRequest(client)
}


func (client *Client) SetResponse (response *Response) {
    client.ServerResponses = append(client.ServerResponses, response)
}


func CreateClient (id int) *Client {
    
    return &Client{
        Id: id,
        ServerResponses: make([]*Response, 0),
    }
}