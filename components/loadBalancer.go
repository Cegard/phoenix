package components

import (
    "phoenix/utils"
    _"fmt"
)

var loadBalancerInstance *loadBalancer


type loadBalancer struct {
    Balancer
    services []*service
}


func buildComparerById (id uint) func(*interface{}) bool {
    
    
    return func (item *interface{}) bool {
        
        return &item.Id == id
    }
}


func buildComparerByLoad () func(*interface{}) bool {
    
    
    return func (item *interface{}) bool {
        
        return item.currentLoad.Value < utils.MAX_SERVICE_CAPACITY
    }
}


func GetLoadBalancer() *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make([]*service, 0, 0),
        }
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) AddService (service *service) {
    balancerInstance.services = append(balancerInstance.services, service)
}


func (balancerInstance *loadBalancer) RemoveService (serviceId uint) {
    var index, wasIndexFound = utils.FindIndexBy(
        balancerInstance.services,
        buildComparerById(serviceId),
    )
    
    if wasIndexFound {
        balancerInstance.services = append(
            balancerInstance.services[ : index],
            balancerInstance.services[index + 1 : ]...
        )
    }
}


func (balancerInstance *loadBalancer) getNextFreeServer() *service {
    var index, wasIndexFound = utils.FindIndexBy(balancerInstance.services, buildComparerByLoad())
    
    if wasIndexFound {
        
        return balancerInstance.services[index]
    }
    
    return nil
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    server := balancerInstance.getNextFreeServer()
    
    if server == nil {
        server := NewService()
        balancerInstance.AddService(server)
        server.AddRequest(clientRequest)
    }
    
    server.AddRequest(clientRequest)
}