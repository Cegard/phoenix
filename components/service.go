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
    channel *chan int
    syncGroup *sync.WaitGroup
}


func (server *service) processRequest (clientRequest *request) {
    time.Sleep(time.Second * clientRequest.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        clientRequest.SetStatus(utils.SUCCEEDED_STATUS)
    } else {
        clientRequest.SetStatus(utils.FAILED_STATUS)
    }
    
    server.currentLoad.DecreaseLoad()
    server.syncGroup.Done()
    <- *server.channel
}


func (server *service) AddRequest (clientRequest *request) {
    *server.channel <- 1
    server.currentLoad.IncreaseLoad()
    server.syncGroup.Add(1)
    go server.processRequest(clientRequest)
}


func NewService(wg *sync.WaitGroup) *service {
    currentCount++
    var channel = make(chan int, utils.MAX_SERVICE_CAPACITY)
    
    return &service {
        Id: currentCount,
        currentLoad: &Load{Value: 0},
        channel: &channel,
        syncGroup: wg,
    }
}
