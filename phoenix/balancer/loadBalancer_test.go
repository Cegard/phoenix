package balancer

import (
    "phoenix/client"
    "phoenix/utils"
    "testing"
    "fmt"
    "time"
)


func TestCreateLoadBalancer (t *testing.T) {
    
    if GetLoadBalancer() == nil {
        t.Errorf("Balancer instance is nil\n")
    }
}


func TestGetLoadBalancer (t *testing.T) {
    var loadBalancer1 = GetLoadBalancer()
    var loadBalancer2 = GetLoadBalancer()
    
    if fmt.Sprintf("%p", loadBalancer1) != fmt.Sprintf("%p", loadBalancer2) {
        t.Errorf("components.GetLoadBalancer creates a new object instance\n")
    }
}


func TestServerInstanceCreation (t *testing.T) {
    
    if GetLoadBalancer().TotalRunningInstances() < utils.MinRunningServices {
        t.Errorf("Live services doesn't match min requirements\n")
    }
}


func TestServerInstanceDynamicCreation (t *testing.T) {
    var client = client.NewClient(0)
    
    for i := 0; i < (utils.MinRunningServices * utils.MaxServiceCapacity * 2); i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    if GetLoadBalancer().TotalRunningInstances() != utils.MinRunningServices * 2 {
        t.Errorf("Load Balancer instance not scaling up\n")
    }
}


func TestServerInstanceDynamicRemoval (t *testing.T) {
    var client = client.NewClient(0)
    
    for i := 0; i < (utils.MinRunningServices * utils.MaxServiceCapacity * 2); i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    var firstServicesCount = GetLoadBalancer().TotalRunningInstances()
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    var secondServicesCount = GetLoadBalancer().TotalRunningInstances()
    
    if secondServicesCount >= firstServicesCount {
        t.Errorf("Load Balancer instance not scaling down\n")
    }
}
