package components

var loadBalancerInstance *loadBalancer


type Balancer interface {
    AddService (*service)
    RemoveService (int)
}


type loadBalancer struct {
    services map[int] *service
}


func GetLoadBalancer() *loadBalancer {
    
    if loadBalancerInstance == nil {
        loadBalancerInstance = &loadBalancer {
            services: make(map[int] *service),
        }
    }
    
    return loadBalancerInstance
}


func (balancerInstance *loadBalancer) AddService (service *service) {
    balancerInstance.services[service.Id] = service
}


func (balancerInstance *loadBalancer) RemoveService (serviceId int) {
    delete(balancerInstance.services, serviceId)
}