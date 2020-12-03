package ui

import (
    "phoenix/client"
    "phoenix/balancer"
    "fmt"
    "strconv"
    "strings"
)


func sendRequests(client *client.Client, requestsNumber int) {
    
    for index := 0; index < requestsNumber ; index++ {
        
        balancer.GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
}


func ProcessUserCommands(client *client.Client, command string) {
    var input = strings.Split(command, " ")
    
    switch input[0] {
        
        case "send":
            requsetsNumbers, _ := strconv.Atoi(input[1])
            sendRequests(client, requsetsNumbers)
        
        case "serverStatus":
            fmt.Printf("\nProcessed requests so far: %d\n", len(client.ServerResponses))
            balancer.GetLoadBalancer().PrintStatus()
        
        case "serviceStatus":
            serviceId, _ := strconv.Atoi(input[1])
            balancer.GetLoadBalancer().PrintServiceStatus(serviceId)
        
        default:
            fmt.Printf("Command not recognized\n")
    }
}
