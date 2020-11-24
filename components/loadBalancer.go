package components

import (
    "phoenix/utils"
    "golang.org/x/tools/go/types/typeutil"
)

var loadBalancerInstance *loadBalancer


type loadBalancer struct {
    services *typeutil.Map
}


func GetLoadBalancer() *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: new(typeutil.Map),
        }
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) AddService (service *service) {
    balancerInstance.services.Set(service.Id, service)
}


func (balancerInstance *loadBalancer) RemoveService (serviceId ServiceId) {
    balancerInstance.services.Delete(serviceId)
}


func (balancerInstance *loadBalancer) getNextFreeServer() *service {
    var index = 0
    var freeService *service = nil
    
    for index < balancerInstance.services.Len() && freeService != nil {
                
        if balancerInstance.
           services.At(balancerInstance.services.Keys()[index]).(*service).CurrentLoad <
                utils.MAX_SERVICE_CAPACITY {
            freeService = balancerInstance.
                          services.At(balancerInstance.services.Keys()[index]).(*service)
        } else {
            index++
        }
    }
    
    return freeService
}


func (balancerInstance *loadBalancer) AssignRequest (clientRequest *request) {
    server := balancerInstance.getNextFreeServer()
    
    if server != nil {
        server := NewService()
        balancerInstance.AddService(server)
        server.AddRequest(clientRequest)
    }
    
    server.AddRequest(clientRequest)
}