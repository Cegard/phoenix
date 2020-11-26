package components

import (
    "phoenix/utils"
    "time"
    "fmt"
)


type request struct {
    TimeToProcess time.Duration
    Client *Client
    Status int
}


func NewRequest(client *Client) *request {
    
    return &request {
        Client: client,
        TimeToProcess: time.Duration(utils.RandomInt(
            utils.MIN_PROCESS_TIME,
            utils.MAX_PROCESS_TIME),
        ),
        Status: utils.WAITING_STATUS,
    }
}


func (clientRequest *request) SetStatus (processResult int) {
    fmt.Printf("Request with code %d finished after %d seconds\n", processResult, clientRequest.TimeToProcess)
    clientRequest.Status = processResult
}
