package components

import (
    "phoenix/utils"
    "sync"
    "time"
)

var loadBalancerInstance *loadBalancer


type loadBalancer struct {
    Balancer
    sync.Mutex
    services map[int] *service
    syncGroup *sync.WaitGroup
}


func fillLoadBalancer (balancerInstance *loadBalancer) {
    balancerInstance.Lock()
    
    for i := 0; i < utils.MIN_RUNING_SERVICES; i++ {
        loadBalancerInstance.addService(CreateService(balancerInstance.syncGroup))
    }
    
    balancerInstance.Unlock()
}


func GetLoadBalancer(wg *sync.WaitGroup) *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make(map[int] *service),
            syncGroup: wg,
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
        var server = CreateService(balancerInstance.syncGroup)
        serverId = server.Id
        balancerInstance.addService(server)
    }
    
    return balancerInstance.services[serverId].AddRequest(clientRequest)
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    
    if !balancerInstance.assignRequest(clientRequest) {
        balancerInstance.AssignRequest(clientRequest)
    }
}
