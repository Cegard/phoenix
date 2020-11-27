package components

import (
    "phoenix/utils"
    "time"
)


type request struct {
    TimeToProcess time.Duration
    Client *Client
}


func NewRequest(client *Client) *request {
    
    return &request {
        Client: client,
        TimeToProcess: time.Duration(utils.RandomInt(
            utils.MIN_PROCESS_TIME,
            utils.MAX_PROCESS_TIME),
        ),
    }
}
