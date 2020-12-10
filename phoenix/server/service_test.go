package server

import (
    "phoenix/info"
    "phoenix/client"
    "phoenix/utils"
    "fmt"
    "sync"
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestProcessRequest (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    var counter = &utils.Counter{}
    
    wg.Add(1)
    service.processRequest(client.MakeRequest(), info.NewInfo(), counter)
    
    assert.NotEqual(
        t,
        client.GetResponses(),
        0,
        "Service is not processing requests",
    )
    assert.Equal(
        t,
        1,
        counter.GetCount(),
        "Service is not counting processed requests",
    )
}


func TestServerLoadIncreases (t *testing.T) {
    var client = client.NewClient(0)
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    wg.Add(1)
    service.AddRequest(client.MakeRequest(), info.NewInfo(), &utils.Counter{})
    
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
    service.AddRequest(client.MakeRequest(), info.NewInfo(), &utils.Counter{})
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
    service.AddRequest(client.MakeRequest(), info.NewInfo(), &utils.Counter{})
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    assert.GreaterOrEqual(
        t,
        service.successRequests.GetCount() + service.failedRequests.GetCount(),
        1,
        "Service is not updating it's history",
    )
}


func TestAddRequest (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    var client = client.NewClient(0)
    var info = info.NewInfo()
    
    assert.True(
        t,
        service.AddRequest(client.MakeRequest(), info, &utils.Counter{}),
        "Service is not adding new requests",
    )
    
    for i := 0; i < utils.MaxServiceCapacity; i++ {
        service.AddRequest(client.MakeRequest(), info, &utils.Counter{})
    }
    
    assert.False(
        t,
        service.AddRequest(client.MakeRequest(), info, &utils.Counter{}),
        "Service is not rejecting new requests when has no room",
    )
}


func TestHasRoom (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    var client = client.NewClient(0)
    var info = info.NewInfo()
    
    assert.True(
        t,
        service.HasRoom(),
        "Service should have room",
    )
    
    for i := 0; i < utils.MaxServiceCapacity; i++ {
        service.AddRequest(client.MakeRequest(), info, &utils.Counter{})
    }
    
    assert.False(
        t,
        service.HasRoom(),
        "Service shouldn't have room",
    )
}


func TestIsIdle (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    var client = client.NewClient(0)
    
    service.AddRequest(client.MakeRequest(), info.NewInfo(), &utils.Counter{})
    
    assert.False(
        t,
        service.IsIdle(),
        "Service is not on non idle state",
    )
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    assert.True(
        t,
        service.IsIdle(),
        "Service is not on idle state",
    )
}


func TestIsUp (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    assert.True(
        t,
        service.IsUp(),
        "Service is not going up",
    )
}


func TestShutDown (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    service.ShutDown()
    assert.False(
        t,
        service.IsUp(),
        "Service is not shutting down",
    )
}


func TestString (t *testing.T) {
    var wg sync.WaitGroup
    var service = NewService(0, &wg)
    
    assert.Equal(
        t,
        fmt.Sprintf(
            "Service: %d\n -- Currently processing: %d\n -- Is running: %t\n" +
                    " -- Total processed requests: %d\n -- Succeeded: %d\n -- Failed: %d\n" +
                    " -- Success rate: %f\n\n",
            0,
            0,
            true,
            0,
            0,
            0,
            float64(1),
        ),
        fmt.Sprintf("%s", service),
        "Service is not printing well",
    )
}
