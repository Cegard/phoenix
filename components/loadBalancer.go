package components

import (
    "phoenix/utils"
    "sort"
    "sync"
)

var loadBalancerInstance *loadBalancer


type loadBalancer struct {
    Balancer
    services []*service
    syncGroup *sync.WaitGroup
}


func GetLoadBalancer(wg *sync.WaitGroup) *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make([]*service, 0, 0),
            syncGroup: wg,
        }
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) AddService (service *service) {
    balancerInstance.services = append(balancerInstance.services, service)
}


func (balancerInstance *loadBalancer) RemoveService (serviceId uint) {
    var index = sort.Search(
        -1,
        func (i int) bool {
            
            return balancerInstance.services[i].Id == serviceId
        },
    )
    
    if index >= 0 && len(balancerInstance.services) > 0 {
        balancerInstance.services = append(
            balancerInstance.services[ : index],
            balancerInstance.services[index + 1 : ]...
        )
    }
}


func (balancerInstance *loadBalancer) getNextFreeServer() *service {
    var index = sort.Search(
        -1,
        func (i int) bool {
            
            return balancerInstance.services[i].currentLoad.Value < utils.MAX_SERVICE_CAPACITY
        },
    )
    
    if index >= 0 && len(balancerInstance.services) > 0 {
        return balancerInstance.services[index]
    }
    
    return nil
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    server := balancerInstance.getNextFreeServer()
    
    if server == nil {
        server := NewService(balancerInstance.syncGroup)
        balancerInstance.AddService(server)
        server.AddRequest(clientRequest)
    } else {
        server.AddRequest(clientRequest)
    }
}