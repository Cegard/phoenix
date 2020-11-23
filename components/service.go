package components

import (
    "time"
    "sync"
    "phoenix/utils"
)


type service struct {
    sync.Mutex
    Id int
    CurrentLoad int
    channel chan int
}


func (server *service) processRequest (clientRequest *request) {
    server.Lock()
    server.CurrentLoad++
    time.Sleep(time.Second * clientRequest.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        clientRequest.SetStatus(utils.SUCCEEDED_STATUS)
    } else {
        clientRequest.SetStatus(utils.FAILED_STATUS)
    }
    
    server.Unlock()
    server.CurrentLoad--
    <- server.channel
}


func (server *service) AddRequest (clientRequest *request) {
    server.channel <- 1
    go server.processRequest(clientRequest)
}


func (server *service) NewService (id int) *service {
    return &service {
        Id: id,
        CurrentLoad: 0,
        channel: make(chan int, utils.MAX_SERVICE_CAPACITY),
    }
}
