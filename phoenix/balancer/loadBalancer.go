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
    syncGroup *sync.WaitGroup
    servicesIdentifier *utils.Counter
    services map[int] *server.Service
    stats *Info
}


func GetLoadBalancer() *LoadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &LoadBalancer {
            services: make(map[int] *server.Service),
            syncGroup: &sync.WaitGroup{},
            servicesIdentifier: &utils.Counter{},
            stats: NewInfo(),
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
        var server = server.NewService(balancer.servicesIdentifier.GetCount(), balancer.syncGroup)
        balancer.services[server.Id] = server
    }
    
    balancer.Unlock()
}


func (balancer *LoadBalancer) removeIdleServices () {
    loadBalancerInstance.Lock()
    
    for serviceId := range balancer.services {
        
        if balancer.services[serviceId].IsIdle() &&
                len(balancer.services) > utils.MinRunningServices {
            balancer.services[serviceId].ShutDown()
            balancer.stats.RegisterStat(
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
        var server = server.NewService(balancer.servicesIdentifier.GetCount(), balancer.syncGroup)
        serverId = server.Id
        balancer.services[server.Id] = server
    }
    
    return balancer.services[serverId].AddRequest(request, balancer.stats.RegisterStat)
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
    
    return balancer.
           stats.
           GetAllInfo(fmt.Sprintf("\n\nTotal running services: %d\n\n", len(balancer.services)))
}


func (balancer *LoadBalancer) GetServiceStatus (serviceId int) string {
    balancer.Lock()
    defer balancer.Unlock()
    
    return balancer.stats.GetEntryInfo(serviceId)
}


func (balancer *LoadBalancer) TotalRunningInstances() int {
    balancer.Lock()
    defer balancer.Unlock()
    
    return len(balancer.services)
}
