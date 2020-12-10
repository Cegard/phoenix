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


func ProcessUserCommands (client *client.Client, command string) (string, error) {
    var input = strings.Split(command, " ")
    
    switch input[0] {
        
        case fmt.Sprintf("%s", utils.Send):
            requestsNumber, err := strconv.Atoi(input[1])
            
            if err == nil {
                sendRequests(client, requestsNumber)
                
                return "", nil
            } else {
                
                return utils.CouldNotProcessMsg, utils.NewNotNumberError(input[1])
            }
        
        case fmt.Sprintf("%s", utils.ServerStatus):
            
            return fmt.Sprintf(
                "\nProcessed requests so far: %d\n, %s",
                len(client.GetResponses()),
                joinStats(balancer.GetLoadBalancer().GetStatus()),
            ), nil
        
        case fmt.Sprintf("%s", utils.ServiceStatus):
            serviceId, err := strconv.Atoi(input[1])
            
            if err == nil {
                
                return balancer.GetLoadBalancer().Stats.GetEntryInfo(serviceId), nil
            } else {
                
                return utils.CouldNotProcessMsg, utils.NewNotNumberError(input[1])
            }
        
        default:
            
            return utils.NotRecognizedMsg, nil
    }
}
