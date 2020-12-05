package ui

import (
    "phoenix/client"
    "phoenix/balancer"
    "phoenix/utils"
    "fmt"
    "strconv"
    "strings"
)


func sendRequests(client *client.Client, requestsNumber int) {
    
    for index := 0; index < requestsNumber ; index++ {
        
        balancer.GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
}


func ProcessUserCommands (client *client.Client, command string) (string, error) {
    var input = strings.Split(command, " ")
    
    switch input[0] {
        
        case "send":
            requestsNumber, err := strconv.Atoi(input[1])
            
            if err == nil {
                sendRequests(client, requestsNumber)
                
                return "", nil
            } else {
                
                return "", utils.NewNotNumberError(input[1])
            }
        
        case "serverStatus":
            
            return fmt.Sprintf(
                "\nProcessed requests so far: %d\n, %s",
                len(client.ServerResponses,
            ), balancer.GetLoadBalancer().GetStatus()), nil
        
        case "serviceStatus":
            serviceId, err := strconv.Atoi(input[1])
            
            if err == nil {
                
                return balancer.GetLoadBalancer().GetServiceStatus(serviceId), nil
            } else {
                
                return "", utils.NewNotNumberError(input[1])
            }
        
        default:
            
            return fmt.Sprintf("Command not recognized\n"), nil
    }
}
