package components

import (
    "phoenix/utils"
    "sync"
    "time"
    "fmt"
)

var loadBalancerInstance *loadBalancer


type loadBalancer struct {
    sync.Mutex
    SyncGroup *sync.WaitGroup
    services map[int] *service
}


func fillLoadBalancer (balancerInstance *loadBalancer) {
    balancerInstance.Lock()
    
    for i := 0; i < utils.MIN_RUNING_SERVICES; i++ {
        loadBalancerInstance.addService(CreateService(balancerInstance.SyncGroup))
    }
    
    balancerInstance.Unlock()
}


func GetLoadBalancer() *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make(map[int] *service),
            SyncGroup: &sync.WaitGroup{},
        }
        
        fillLoadBalancer(loadBalancerInstance)
        go loadBalancerInstance.removeServicesContinously()
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) addService (server *service) {
    balancerInstance.services[server.Id] = server
}


func (balancerInstance *loadBalancer) removeIdleServices () {
    loadBalancerInstance.Lock()
    
    for serviceId := range balancerInstance.services {
        
        if balancerInstance.services[serviceId].IsIdle() &&
                len(balancerInstance.services) > utils.MIN_RUNING_SERVICES {
            delete(balancerInstance.services, serviceId)
        }
    }
    
    loadBalancerInstance.Unlock()
}


func (balancerInstance *loadBalancer) removeServicesContinously () {
    
    for true {
        time.Sleep(time.Second)
        balancerInstance.removeIdleServices()
    }
}


func (balancerInstance *loadBalancer) getNextFreeServerId() (int, bool) {
    
    for index := range balancerInstance.services {
        
        if balancerInstance.services[index].HasRoom() {
            
            return index, true
        }
    }
    
    return -1, false
}


func (balancerInstance *loadBalancer) assignRequest (clientRequest *request) bool {
    defer balancerInstance.Unlock()
    balancerInstance.Lock()
    var serverId, wasFound = balancerInstance.getNextFreeServerId()
    
    if !wasFound {
        var server = CreateService(balancerInstance.SyncGroup)
        serverId = server.Id
        balancerInstance.addService(server)
    }
    
    return balancerInstance.services[serverId].AddRequest(clientRequest)
}


func (balancerInstance *loadBalancer) tryToAssignRequest (clientRequest *request) {
    
    if !balancerInstance.assignRequest(clientRequest) {
        balancerInstance.AssignRequest(clientRequest)
    }
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    go balancerInstance.tryToAssignRequest(clientRequest)
}


func (balancerInstance *loadBalancer) PrintStatus() {
    fmt.Printf("\n\nTotal running services: %d\n\n", len(balancerInstance.services))
    
    for index := range balancerInstance.services {
        fmt.Printf("%s", balancerInstance.services[index])
    }
}


func (balancerInstance *loadBalancer) PrintServiceStatus (serviceId int) {
    fmt.Printf("%s", balancerInstance.services[serviceId])
}
