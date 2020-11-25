package components

import (
    "time"
    "phoenix/utils"
)

var currentCount uint = 0


type service struct {
    Id uint
    currentLoad *Load
    channel chan int
}


func (server *service) processRequest (clientRequest *request) {
    time.Sleep(time.Second * clientRequest.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        clientRequest.SetStatus(utils.SUCCEEDED_STATUS)
    } else {
        clientRequest.SetStatus(utils.FAILED_STATUS)
    }
    
    server.currentLoad.DecreaseLoad()
    <- server.channel
}


func (server *service) AddRequest (clientRequest *request) {
    server.channel <- 1
    server.currentLoad.IncreaseLoad()
    go server.processRequest(clientRequest)
}


func NewService() *service {
    currentCount++
    
    return &service {
        Id: currentCount,
        currentLoad: &Load{Value: 0},
        channel: make(chan int, utils.MAX_SERVICE_CAPACITY),
    }
}
