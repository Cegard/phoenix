package components

import (
    "fmt"
    "strconv"
    "strings"
)


func sendRequests(client *Client, requestsNumber int) {
    
    for index := 0; index < requestsNumber ; index++ {
        
        GetLoadBalancer().AssignRequest(client.MakeRequest())
    }
}


func ProcessUserCommands(client *Client, command string) {
    input := strings.Split(command, " ")
    
    switch input[0] {
        
        case "send":
            requsetsNumbers, _ := strconv.Atoi(input[1])
            sendRequests(client, requsetsNumbers)
        
        case "serverStatus":
            fmt.Printf("\nProcessed requests so far: %d\n", len(client.ServerResponses))
            GetLoadBalancer().PrintStatus()
        
        case "serviceStatus":
            serviceId, _ := strconv.Atoi(input[1])
            GetLoadBalancer().PrintServiceStatus(serviceId)
        
        default:
            fmt.Printf("Command not recognized\n")
    }
}
