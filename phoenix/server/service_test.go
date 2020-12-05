package server

import (
    "phoenix/client"
    "phoenix/utils"
    "testing"
    "sync"
    "time"
)

var dummyRegister = func (i int, s string){}


func TestProcessRequest (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.processRequest(client.MakeRequest(), dummyRegister)
    
    if len(client.ServerResponses) == 0 {
        t.Errorf("Service is not processing requests")
    }
}


func TestServerLoadIncreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    
    if service.currentCount.GetCount() < 1 {
        t.Errorf("Service is not increasing it's load")
    }
}


func TestServerLoadDecreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    previousProcessingLoad := service.currentCount.GetCount()
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    if service.currentCount.GetCount() >= previousProcessingLoad {
        t.Errorf("Service is not decreasing it's load")
    }
}


func TestServerUpdatesHistory (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), dummyRegister)
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    if service.successRequests.GetCount() + service.failedRequests.GetCount() < 1 {
        t.Errorf("Service is not updating it's history")
    }
}
