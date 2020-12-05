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
    sync.Mutex
    Id int
    IsRunning bool
    currentLoad *Load
    syncGroup *sync.WaitGroup
    successRequests *Load
    failedRequests *Load
}


func NewService(mainWaitGroup *sync.WaitGroup) *Service {
    currentCount++
    
    return &Service {
        Id: currentCount,
        IsRunning: true,
        currentLoad: &Load{},
        syncGroup: mainWaitGroup,
        successRequests: &Load{},
        failedRequests: &Load{},
    }
}


func (server *Service) processRequest (request *messages.Request, register func(int, string)) {
    time.Sleep(time.Second * request.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SuccessProbability {
        server.successRequests.IncreaseLoad()
        request.RespondTo(messages.NewResponse(utils.SucceededStatus, server.Id))
    } else {
        server.failedRequests.IncreaseLoad()
        request.RespondTo(messages.NewResponse(utils.FailedStatus, server.Id))
    }
    
    server.currentLoad.DecreaseLoad()
    register(server.Id, fmt.Sprintf("%s", server))
    server.syncGroup.Done()
}


func (server *Service) AddRequest (request *messages.Request, register func(int, string)) bool {
    
    if server.HasRoom() {
        server.currentLoad.IncreaseLoad()
        register(server.Id, fmt.Sprintf("%s", server))
        server.syncGroup.Add(1)
        go server.processRequest(request, register)
        
        return true
    }
    
    return false
}


func (server *Service) HasRoom() bool {
    server.Lock()
    defer server.Unlock()
    
    return server.currentLoad.GetValue() < utils.MaxServiceCapacity
}


func (server *Service) IsIdle() bool {
    server.Lock()
    defer server.Unlock()
    
    return server.currentLoad.GetValue() == 0
}


func (server *Service) String() string {
    server.Lock()
    defer server.Unlock()
    var successRate = 0.0
    var totalRequests = (server.successRequests.GetValue() + server.failedRequests.GetValue())
    
    if totalRequests == 0 {
        successRate = 1.0
    } else {
        successRate = float64(server.successRequests.GetValue()) / float64(totalRequests)
    }
    
    return fmt.Sprintf(
        "Service: %d\n -- Currently processing: %d\n -- Is running: %t\n" +
                " -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n" +
                " -- Success rate: %f\n\n",
        server.Id,
        server.currentLoad.GetValue(),
        server.IsRunning,
        server.successRequests.GetValue() + server.failedRequests.GetValue(),
        server.successRequests.GetValue(),
        server.failedRequests.GetValue(),
        successRate,
    )
}


func (server *Service) ShutDown() {
    server.Lock()
    server.IsRunning = false
    server.Unlock()
}
