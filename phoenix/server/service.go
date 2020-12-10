package server

import (
    "phoenix/info"
    "phoenix/messages"
    "phoenix/utils"
    "sync"
    "time"
    "fmt"
)


type Service struct {
    sync.Mutex
    Id int
    isRunning bool
    currentCount *utils.Counter
    syncGroup *sync.WaitGroup
    successRequests *utils.Counter
    failedRequests *utils.Counter
}


func NewService(id int, mainWaitGroup *sync.WaitGroup) *Service {
    
    return &Service {
        Id: id,
        isRunning: true,
        currentCount: &utils.Counter{},
        syncGroup: mainWaitGroup,
        successRequests: &utils.Counter{},
        failedRequests: &utils.Counter{},
    }
}


func (server *Service) processRequest (
        request *messages.Request,
        registerer *info.Info,
        counter *utils.Counter,
    ) {
    time.Sleep(time.Second * request.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SuccessProbability {
        server.successRequests.IncreaseCount()
        request.RespondTo(messages.NewResponse(utils.SucceededStatus, server.Id))
    } else {
        server.failedRequests.IncreaseCount()
        request.RespondTo(messages.NewResponse(utils.FailedStatus, server.Id))
    }
    
    server.currentCount.DecreaseCount()
    registerer.RegisterStat(server.Id, fmt.Sprintf("%s", server))
    counter.IncreaseCount()
    server.syncGroup.Done()
}


func (server *Service) AddRequest (
        request *messages.Request,
        registerer *info.Info,
        counter *utils.Counter,
    ) bool {
    
    if server.HasRoom() && server.IsUp() {
        server.currentCount.IncreaseCount()
        registerer.RegisterStat(server.Id, fmt.Sprintf("%s", server))
        server.syncGroup.Add(1)
        go server.processRequest(request, registerer, counter)
        
        return true
    }
    
    return false
}


func (server *Service) HasRoom() bool {
    server.Lock()
    defer server.Unlock()
    
    return server.currentCount.GetCount() < utils.MaxServiceCapacity
}


func (server *Service) IsIdle() bool {
    server.Lock()
    defer server.Unlock()
    
    return server.currentCount.GetCount() == 0
}


func (server *Service) String() string {
    server.Lock()
    defer server.Unlock()
    var successRate = 0.0
    var totalRequests = (server.successRequests.GetCount() + server.failedRequests.GetCount())
    
    if totalRequests == 0 {
        successRate = 1.0
    } else {
        successRate = float64(server.successRequests.GetCount()) / float64(totalRequests)
    }
    
    return fmt.Sprintf(
        "Service: %d\n -- Currently processing: %d\n -- Is running: %t\n" +
                " -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n" +
                " -- Success rate: %f\n\n",
        server.Id,
        server.currentCount.GetCount(),
        server.isRunning,
        server.successRequests.GetCount() + server.failedRequests.GetCount(),
        server.successRequests.GetCount(),
        server.failedRequests.GetCount(),
        successRate,
    )
}


func (server *Service) ShutDown() {
    server.Lock()
    server.isRunning = false
    server.Unlock()
}


func (server *Service) IsUp() bool {
    server.Lock()
    defer server.Unlock()
    
    return server.isRunning
}
