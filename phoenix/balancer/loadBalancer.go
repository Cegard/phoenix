package balancer

import (
    "phoenix/info"
    "phoenix/messages"
    "phoenix/server"
    "phoenix/utils"
    "sync"
    "time"
    "fmt"
)

var loadBalancerInstance *LoadBalancer


type LoadBalancer struct {
    sync.Mutex
    SyncGroup *sync.WaitGroup
    Stats *info.Info
    servicesIdentifier *utils.Counter
    services map[int] *server.Service
    requestsCounter *utils.Counter
}


func GetLoadBalancer() *LoadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &LoadBalancer {
            services: make(map[int] *server.Service),
            SyncGroup: &sync.WaitGroup{},
            servicesIdentifier: &utils.Counter{},
            Stats: info.NewInfo(),
            requestsCounter: &utils.Counter{},
        }
        
        fillLoadBalancer(loadBalancerInstance)
        go loadBalancerInstance.removeServicesContinuously()
    }
    
    return loadBalancerInstance
}


func fillLoadBalancer (balancer *LoadBalancer) {
    balancer.Lock()
    
    for id := 0; id < utils.MinRunningServices; id++ {
        balancer.servicesIdentifier.IncreaseCount()
        var server = server.NewService(balancer.servicesIdentifier.GetCount(), balancer.SyncGroup)
        balancer.services[server.Id] = server
        balancer.Stats.RegisterStat(server.Id, fmt.Sprintf("%s", server))
    }
    
    balancer.Unlock()
}


func (balancer *LoadBalancer) removeIdleServices () {
    loadBalancerInstance.Lock()
    
    for serviceId := range balancer.services {
        
        if balancer.services[serviceId].IsIdle() &&
                len(balancer.services) > utils.MinRunningServices {
            balancer.services[serviceId].ShutDown()
            balancer.Stats.RegisterStat(
                serviceId,
                fmt.Sprintf("%s", balancer.services[serviceId]),
            )
            delete(balancer.services, serviceId)
        }
    }
    
    loadBalancerInstance.Unlock()
}


func (balancer *LoadBalancer) removeServicesContinuously () {
    
    for true {
        time.Sleep(time.Second)
        balancer.removeIdleServices()
    }
}


func (balancer *LoadBalancer) getNextFreeServerId() (int, bool) {
    
    for index := range balancer.services {
        
        if balancer.services[index].HasRoom() {
            
            return index, true
        }
    }
    
    return -1, false
}


func (balancer *LoadBalancer) assignRequest (request *messages.Request) bool {
    balancer.Lock()
    defer balancer.Unlock()
    var serverId, wasFound = balancer.getNextFreeServerId()
    
    if !wasFound {
        balancer.servicesIdentifier.IncreaseCount()
        var server = server.NewService(balancer.servicesIdentifier.GetCount(), balancer.SyncGroup)
        serverId = server.Id
        balancer.services[server.Id] = server
    }
    
    return balancer.
           services[serverId].
           AddRequest(request, balancer.Stats, balancer.requestsCounter)
}


func (balancer *LoadBalancer) tryToAssignRequest (request *messages.Request) {
    
    if !balancer.assignRequest(request) {
        balancer.AssignRequest(request)
    }
}


func (balancer *LoadBalancer) AssignRequest (request *messages.Request) {
    go balancer.tryToAssignRequest(request)
}


func (balancer *LoadBalancer) GetStatus() []string {
    balancer.Lock()
    defer balancer.Unlock()
    var stats = []string{
        fmt.Sprintf("\n\nTotal running services: %d\n\n",len(balancer.services)),
        fmt.Sprintf("Processed requests so far: %d\n\n", balancer.requestsCounter.GetCount()),
    }
    
    return append(
        stats,
        balancer.Stats.GetAllInfo()...,
    )
}


func (balancer *LoadBalancer) TotalRunningInstances() int {
    balancer.Lock()
    defer balancer.Unlock()
    
    return len(balancer.services)
}
