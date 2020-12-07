package balancer

import (
    "phoenix/client"
    "phoenix/utils"
    "fmt"
    "time"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestCreateLoadBalancer (t *testing.T) {
    assert.NotEqual(t, GetLoadBalancer(), nil, "Balancer instance is nil\n")
}


func TestGetLoadBalancer (t *testing.T) {
    var loadBalancer1 = GetLoadBalancer()
    var loadBalancer2 = GetLoadBalancer()
    
    assert.Equal(
        t,
        fmt.Sprintf("%p", loadBalancer1),
        fmt.Sprintf("%p", loadBalancer2),
        "components.GetLoadBalancer creates a new object instance\n",
    )
}


func TestBalancerInstancesCreation (t *testing.T) {
    assert.Equal(
        t,
        GetLoadBalancer().TotalRunningInstances(),
        utils.MinRunningServices,
        "Live services doesn't match min requirements\n",
    )
}


func TestBalancerInstancesDynamicCreation (t *testing.T) {
    var client = client.NewClient(0)
    
    for i := 0; i < (utils.MinRunningServices * utils.MaxServiceCapacity * 2); i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    assert.Equal(
        t,
        GetLoadBalancer().TotalRunningInstances(),
        utils.MinRunningServices * 2,
        "Load Balancer instance not scaling up\n",
    )
}


func TestBalancerInstancesDynamicRemoval (t *testing.T) {
    var client = client.NewClient(0)
    
    for i := 0; i < (utils.MinRunningServices * utils.MaxServiceCapacity * 2); i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    var firstServicesCount = GetLoadBalancer().TotalRunningInstances()
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    var secondServicesCount = GetLoadBalancer().TotalRunningInstances()
    
    assert.Greater(
        t,
        firstServicesCount,
        secondServicesCount,
        "Load Balancer instance not scaling down\n",
    )
}

/*
func TestGetStatus (t *testing.T) {
    
}
*/