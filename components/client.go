package components


type Client struct {
    Id int
}


func (client *Client) MakeRequest () *request {
    
    return NewRequest(client)
}
