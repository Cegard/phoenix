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
    successRequests *utils.Load
    failedRequests *utils.Load
}


func (server *Service) processRequest (request *messages.Request) {
    time.Sleep(time.Second * request.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SuccessProbability {
        server.successRequests.IncreaseLoad()
        request.RespondTo(messages.NewResponse(utils.SucceededStatus, server.Id))
    } else {
        server.failedRequests.IncreaseLoad()
        request.RespondTo(messages.NewResponse(utils.FailedStatus, server.Id))
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
    
    return server.currentLoad.GetValue() < utils.MaxServiceCapacity
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
        successRequests: &utils.Load{},
        failedRequests: &utils.Load{},
    }
}


func (server *Service) String() string {
    
    return fmt.Sprintf(
        "Service: %d\n -- Currently processing: %d\n -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n\n",
        server.Id,
        server.currentLoad.GetValue(),
        server.successRequests.GetValue() + server.failedRequests.GetValue(),
        server.successRequests.GetValue(),
        server.successRequests.GetValue() /
            (server.successRequests.GetValue() + server.failedRequests.GetValue()),
        server.failedRequests.GetValue(),
    )
}
