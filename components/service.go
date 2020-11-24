package components

import (
    "time"
    "sync"
    "phoenix/utils"
    "go/types"
)

type ServiceId uint


func (id ServiceId) String() string {
    return types.TypeString(id, nil)
}


func (id ServiceId) Underlying() types.Type {
    return id
}


var currentCount ServiceId = 0


type service struct {
    sync.Mutex
    Id ServiceId
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


func NewService() *service {
    currentCount++
    
    return &service {
        Id: currentCount,
        CurrentLoad: 0,
        channel: make(chan int, utils.MAX_SERVICE_CAPACITY),
    }
}
