package components

import (
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
    
    if index >= 0 && len(balancerInstance.services) > 1 {
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
            
            return balancerInstance.services[i].HasRoom()
        },
    )
    
    if index >= 0 && len(balancerInstance.services) > 0 {
        return balancerInstance.services[index]
    }
    
    return nil
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    var server = balancerInstance.getNextFreeServer()
    
    if server == nil {
        server = NewService(balancerInstance.syncGroup)
        balancerInstance.AddService(server)
    }
    
    _ = server.AddRequest(clientRequest)
    
}
