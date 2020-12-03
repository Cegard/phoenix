package server

import (
    "phoenix/messages"
    "phoenix/utils"
    "sync"
    "time"
    "fmt"
)

var currentCount int = 0


type Service struct {
    Id int
    currentLoad *utils.Load
    syncGroup *sync.WaitGroup
    successRequests uint
    failedRequests uint
}


func (server *Service) processRequest (request *messages.Request) {
    time.Sleep(time.Second * request.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        server.successRequests++
        request.RespondTo(messages.NewResponse(utils.SUCCEEDED_STATUS, server.Id))
    } else {
        server.failedRequests++
        request.RespondTo(messages.NewResponse(utils.FAILED_STATUS, server.Id))
    }
    
    server.currentLoad.DecreaseLoad()
    server.syncGroup.Done()
}


func (server *Service) AddRequest (request *messages.Request) bool {
    
    if server.HasRoom() {
        server.currentLoad.IncreaseLoad()
        server.syncGroup.Add(1)
        go server.processRequest(request)
        
        return true
    }
    
    return false
}


func (server *Service) HasRoom() bool {
    
    return server.currentLoad.GetValue() < utils.MAX_SERVICE_CAPACITY
}


func (server *Service) IsIdle() bool {
    
    return server.currentLoad.GetValue() == 0
}


func NewService(mainWaitGroup *sync.WaitGroup) *Service {
    currentCount++
    
    return &Service {
        Id: currentCount,
        currentLoad: &utils.Load{},
        syncGroup: mainWaitGroup,
    }
}


func (server *Service) String() string {
    
    return fmt.Sprintf(
        "Service: %d\n -- Currently processing: %d\n -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n\n",
        server.Id,
        server.currentLoad.GetValue(),
        server.successRequests + server.failedRequests,
        server.successRequests,
        server.failedRequests,
    )
}