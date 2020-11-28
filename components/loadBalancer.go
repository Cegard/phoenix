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
    
    for i := 0; i < utils.MIN_RUNING_SERVICES; i++ {
        loadBalancerInstance.AddService(CreateService())
    }
}


func GetLoadBalancer(wg *sync.WaitGroup) *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make(map[int] *service),
            syncGroup: wg,
        }
        
        fillLoadBalancer(loadBalancerInstance)
        go loadBalancerInstance.removeServices()
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) AddService (server *service) {
    loadBalancerInstance.Lock()
    balancerInstance.services[server.Id] = server
    loadBalancerInstance.Unlock()
}


func (balancerInstance *loadBalancer) RemoveService (serviceId int) {
    
    if balancerInstance.services[serviceId].IsIdle() &&
            len(balancerInstance.services) > utils.MIN_RUNING_SERVICES {
        delete(balancerInstance.services, serviceId)
    }
}


func (balancerInstance *loadBalancer) removeServices () {
    
    for true {
        time.Sleep(time.Second)
        loadBalancerInstance.Lock()
        
        for serverId := range balancerInstance.services {
            balancerInstance.RemoveService(serverId)
        }
        
        loadBalancerInstance.Unlock()
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


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    var serverId, wasFound = balancerInstance.getNextFreeServerId()
    
    if !wasFound {
        var server = CreateService()
        serverId = server.Id
        balancerInstance.AddService(server)
    }
    
    balancerInstance.syncGroup.Add(1)
    go balancerInstance.services[serverId].AddRequest(clientRequest, balancerInstance.syncGroup)
}
