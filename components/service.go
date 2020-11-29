package components

import (
    "phoenix/utils"
    "sync"
    "time"
    "fmt"
)

var currentCount int = 0


type service struct {
    Id int
    currentLoad *Load
    syncGroup *sync.WaitGroup
    successRequests uint
    failedRequests uint
}


func (server *service) processRequest (clientRequest *request) {
    time.Sleep(time.Second * clientRequest.TimeToProcess)
    
    if utils.RandomFloat() <= utils.SUCCESS_PROBABILITY {
        server.successRequests++
        clientRequest.Client.SetResponse(CreateResponse(utils.SUCCEEDED_STATUS, server.Id))
    } else {
        server.failedRequests++
        clientRequest.Client.SetResponse(CreateResponse(utils.FAILED_STATUS, server.Id))
    }
    
    server.currentLoad.DecreaseLoad()
    server.syncGroup.Done()
}


func (server *service) AddRequest (clientRequest *request) bool {
    
    if server.HasRoom() {
        server.currentLoad.IncreaseLoad()
        server.syncGroup.Add(1)
        go server.processRequest(clientRequest)
        
        return true
    }
    
    return false
}


func (server *service) HasRoom() bool {
    
    return server.currentLoad.GetValue() < utils.MAX_SERVICE_CAPACITY
}


func (server *service) IsIdle() bool {
    
    return server.currentLoad.GetValue() == 0
}


func CreateService(mainWaitGroup *sync.WaitGroup) *service {
    currentCount++
    
    return &service {
        Id: currentCount,
        currentLoad: new(Load),
        syncGroup: mainWaitGroup,
    }
}


func (server *service) String() string {
    
    return fmt.Sprintf(
        "Service: %d\n -- Currently processing: %d\n -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n\n",
        server.Id,
        server.currentLoad.GetValue(),
        server.successRequests + server.failedRequests,
        server.successRequests,
        server.failedRequests,
    )
}
