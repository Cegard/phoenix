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


func joinStats (stats []string) string {
    var jointStats = stats[0]
    
    for index := 1; index < len(stats); index++ {
        jointStats = fmt.Sprintf("%s\n%s", jointStats, stats[index])
    }
    
    return jointStats
}


func processSendCommand (client *client.Client, requestsString string) (string, error) {
    requestsNumber, err := strconv.Atoi(requestsString)
    
    if err == nil {
        sendRequests(client, requestsNumber)
        
        return "", nil
    } else {
        
        return utils.CouldNotProcessMsg, utils.NewNotNumberError(requestsString)
    }
}


func processServiceStatusCommand (idString string) (string, error) {
    serviceId, err := strconv.Atoi(idString)
    
    if err == nil {
        
        return balancer.GetLoadBalancer().Stats.GetEntryInfo(serviceId), nil
    } else {
        
        return utils.CouldNotProcessMsg, utils.NewNotNumberError(idString)
    }
}


func ProcessUserCommands (client *client.Client, command string) (string, error) {
    var input = strings.Split(command, " ")
    
    switch input[0] {
        
        case fmt.Sprintf("%s", utils.Send):
        
            return processSendCommand(client, input[1])
        
        case fmt.Sprintf("%s", utils.ServerStatus):
            
            return fmt.Sprintf("%s", joinStats(balancer.GetLoadBalancer().GetStatus())), nil
        
        case fmt.Sprintf("%s", utils.ServiceStatus):
            
            return processServiceStatusCommand(input[1])
        
        default:
            
            return utils.NotRecognizedMsg, nil
    }
}
