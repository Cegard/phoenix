package main

import (
    "phoenix/components"
    "fmt"
    "os"
    "bufio"
)


func main() {
    var loadBalancer = components.GetLoadBalancer()
    var client = components.CreateClient(1)
    var hasUserStopped bool
    var in = bufio.NewScanner(os.Stdin)
    
    for !hasUserStopped {
        hasUserStopped = components.ProcessUserCommands(client, loadBalancer, in) 
    }
    
    fmt.Println("Waiting for remaining processes to finish...")
    loadBalancer.SyncGroup.Wait()
}
