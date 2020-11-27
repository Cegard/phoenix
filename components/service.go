package components

import (
    "phoenix/utils"
    "sync"
    "time"
)

var currentCount uint = 0


type service struct {
    Id uint
    currentLoad *Load
    syncGroup *sync.WaitGroup
    channel chan uint
}


func (server *service) processRequest (clientRequest *request) {
    time.Sleep(time.Second * clientRequest.TimeToProcess)
    var statusCode uint
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        statusCode = utils.SUCCEEDED_STATUS
    } else {
        statusCode = utils.FAILED_STATUS
    }
    
    server.currentLoad.DecreaseLoad()
    server.channel <- statusCode
    server.syncGroup.Done()
}


func (server *service) AddRequest (clientRequest *request) uint {
    server.currentLoad.IncreaseLoad()
    server.syncGroup.Add(1)
    go server.processRequest(clientRequest)
    statusCode := <- server.channel
    clientRequest.Client.SetResponse(CreateResponse(statusCode))
    
    return statusCode
}


func (server *service) HasRoom() bool {
    
    return server.currentLoad.Value < utils.MAX_SERVICE_CAPACITY
}


func (server *service) IsIdle() bool {
    
    return server.currentLoad.Value == 0
}


func NewService(wg *sync.WaitGroup) *service {
    currentCount++
    
    return &service {
        Id: currentCount,
        currentLoad: &Load{Value: 0},
        syncGroup: wg,
        channel: make(chan uint, 1),
    }
}
