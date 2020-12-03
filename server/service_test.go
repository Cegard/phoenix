package server

import (
    "phoenix/client"
    "phoenix/utils"
    "testing"
    "sync"
    "time"
)


func TestProcessRequest (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(&wg)
    
    wg.Add(1)
    service.processRequest(client.MakeRequest())
    
    if len(client.ServerResponses) == 0 {
        t.Errorf("Service is not processing requests")
    }
}


func TestServerLoadIncreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(&wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest())
    
    if service.currentLoad.GetValue() < 1 {
        t.Errorf("Service is not increasing it's load")
    }
}


func TestServerLoadDecreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(&wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest())
    previousProcessingLoad := service.currentLoad.GetValue()
    time.Sleep(1 + time.Second * time.Duration(utils.MAX_PROCESS_TIME))
    
    if service.currentLoad.GetValue() >= previousProcessingLoad {
        t.Errorf("Service is not decreasing it's load")
    }
}


func TestServerUpdatesHistory (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(&wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest())
    time.Sleep(1 + time.Second * time.Duration(utils.MAX_PROCESS_TIME))
    
    if service.successRequests + service.failedRequests < 1 {
        t.Errorf("Service is not updating it's history")
    }
}
