package server

import (
    "phoenix/client"
    "phoenix/utils"
    "sync"
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)

var dummyRegister = func (i int, s string){}


func TestProcessRequest (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.processRequest(client.MakeRequest(), dummyRegister)
    
    assert.NotEqual(t, client.ServerResponses, 0, "Service is not processing requests")
}


func TestServerLoadIncreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    
    assert.GreaterOrEqual(
        t,
        service.currentCount.GetCount(),
        1,
        "Service is not increasing it's load",
    )
}


func TestServerLoadDecreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    previousProcessingLoad := service.currentCount.GetCount()
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    assert.Greater(
        t,
        previousProcessingLoad,
        service.currentCount.GetCount(),
        "Service is not decreasing it's load",
    )
}


func TestServerUpdatesHistory (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    assert.GreaterOrEqual(
        t,
        service.successRequests.GetCount() + service.failedRequests.GetCount(),
        1,
        "Service is not updating it's history",
    )
}
