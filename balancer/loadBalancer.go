package balancer

import (
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
    services map[int] *server.Service
}


func fillLoadBalancer (balancer *LoadBalancer) {
    balancer.Lock()
    
    for id := 0; id < utils.MinRunningServices; id++ {
        balancer.services[id] = server.NewService(balancer.SyncGroup)
    }
    
    balancer.Unlock()
}


func GetLoadBalancer() *LoadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &LoadBalancer {
            services: make(map[int] *server.Service),
            SyncGroup: &sync.WaitGroup{},
        }
        
        fillLoadBalancer(loadBalancerInstance)
        go loadBalancerInstance.removeServicesContinuously()
    }
    
    return loadBalancerInstance
}


func (balancer *LoadBalancer) removeIdleServices () {
    loadBalancerInstance.Lock()
    
    for serviceId := range balancer.services {
        
        if balancer.services[serviceId].IsIdle() &&
                len(balancer.services) > utils.MinRunningServices {
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
        var server = server.NewService(balancer.SyncGroup)
        serverId = server.Id
        balancer.services[server.Id] = server
    }
    
    return balancer.services[serverId].AddRequest(request)
}


func (balancer *LoadBalancer) tryToAssignRequest (request *messages.Request) {
    
    if !balancer.assignRequest(request) {
        balancer.AssignRequest(request)
    }
}


func (balancer *LoadBalancer) AssignRequest (request *messages.Request) {
    go balancer.tryToAssignRequest(request)
}


func (balancer *LoadBalancer) GetStatus() string {
    var status = fmt.Sprintf("\n\nTotal running services: %d\n\n", len(balancer.services))
    
    for index := range balancer.services {
        status += fmt.Sprintf("%s", balancer.services[index])
    }
    
    return status
}


func (balancer *LoadBalancer) GetServiceStatus (serviceId int) string {
    
    return fmt.Sprintf("%s", balancer.services[serviceId])
}


func (balancer *LoadBalancer) TotalRunningInstances() int {
    balancer.Lock()
    defer balancer.Unlock()
    
    return len(balancer.services)
}
