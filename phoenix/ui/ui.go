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


func decideOnError(toDoOnOk func() string, err error, newError error) (string, error) {
    
    if err == nil {
        
        return toDoOnOk(), nil
    } else {
        
        return "", newError
    }
}


func wrapSendRequest (client *client.Client, requestsNumber int) func() string {
    
    return func() string {
        sendRequests(client, requestsNumber)
        
        return ""
    }
}


func wrapServiceStatus (serviceId int) func() string {
    
    return func() string {
        
        return balancer.GetLoadBalancer().GetServiceStatus(serviceId)
    }
}


func ProcessUserCommands (client *client.Client, command string) (string, error) {
    var input = strings.Split(command, " ")
    
    switch input[0] {
        
        case "send":
            requestsNumber, err := strconv.Atoi(input[1])
            
            return decideOnError(
                wrapSendRequest(client, requestsNumber),
                err,
                utils.NewNotNumberError(input[1]),
            )
        
        case "serverStatus":
            
            return fmt.Sprintf(
                "\nProcessed requests so far: %d\n, %s",
                len(client.ServerResponses,
            ), balancer.GetLoadBalancer().GetStatus()), nil
        
        case "serviceStatus":
            serviceId, err := strconv.Atoi(input[1])
            
            return decideOnError(
                wrapServiceStatus(serviceId),
                err,
                utils.NewNotNumberError(input[1]),
            )
        
        default:
            
            return fmt.Sprintf("Command not recognized\n"), nil
    }
}
