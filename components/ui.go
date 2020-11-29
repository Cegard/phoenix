package components

import (
    "fmt"
    "bufio"
    "strconv"
    "strings"
)


func sendRequests(client *Client, loadBalancerInstance *loadBalancer, requestsNumber int) {
    
    for index := 0; index < requestsNumber ; index++ {
        
        loadBalancerInstance.AssignRequest(client.MakeRequest())
    }
}


func ProcessUserCommands(client *Client, loadBalancerInstance *loadBalancer, in *bufio.Scanner) bool {
    fmt.Println("Enter a command:")
    fmt.Printf(">: ")
    in.Scan()
    input := strings.Split(in.Text(), " ")
    
    switch input[0] {
        
        case "send":
            requsetsNumbers, _ := strconv.Atoi(input[1])
            sendRequests(client, loadBalancerInstance, requsetsNumbers)
            return false
        
        case "serverStatus":
            fmt.Printf("\nProcessed requests so far: %d\n", len(client.ServerResponses))
            loadBalancerInstance.PrintStatus()
            return false
        
        case "serviceStatus":
            serviceId, _ := strconv.Atoi(input[1])
            loadBalancerInstance.PrintServiceStatus(serviceId)
            return false
        
        case "stop":
            return true
        
        default:
            fmt.Printf("Command not recognized\n")
            return false
    }
}
