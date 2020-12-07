package balancer

import (
    "phoenix/client"
    "phoenix/utils"
    "math"
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


func TestAssignRequest (t *testing.T) {
    var client = client.NewClient(0)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    for i := 0; i < (2 * utils.MaxServiceCapacity); i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    assert.GreaterOrEqual(
        t,
        GetLoadBalancer().TotalRunningInstances(),
        2,
        "Load balancer is not assigning requests",
    )
}


func TestBalancerInstancesCreation (t *testing.T) {
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
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
    
    assert.GreaterOrEqual(
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


func TestTotalRunningInstances (t *testing.T) {
    var requests = 500
    var client = client.NewClient(0)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    for i := 0; i < requests; i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    assert.Equal(
        t,
        int(math.Ceil(float64(requests)/float64(utils.MaxServiceCapacity))),
        GetLoadBalancer().TotalRunningInstances(),
        "The running instances doesn't correspond to the given number",
    )
}


func TestGetServiceStatus (t *testing.T) {
    var client = client.NewClient(0)
    
    GetLoadBalancer().AssignRequest(client.MakeRequest())
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    assert.NotNil(
        t,
        GetLoadBalancer().GetServiceStatus(0),
        "Not retrieving service info",
    )
}


func TestGetStatus (t *testing.T) {
    var asserter = assert.New(t)
    var requests = 500
    var client = client.NewClient(0)
    
    time.Sleep(1 + time.Second * time.Duration(utils.MaxProcessTime))
    
    for i := 0; i < requests; i++ {
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
    
    asserter.Equal(
        fmt.Sprintf(
            "\n\nTotal running services: %d\n\n",
            int(math.Ceil(float64(requests)/float64(utils.MaxServiceCapacity))),
        ),
        GetLoadBalancer().GetStatus()[0],
        "History headers don't corresponds",
    )
    
    asserter.GreaterOrEqual(
        len(GetLoadBalancer().GetStatus()),
        int(math.Ceil(float64(requests)/float64(utils.MaxServiceCapacity))),
        "History entries length doesn't correspond to expected",
    )
}
