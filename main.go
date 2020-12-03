package main

import (
    "phoenix/balancer"
    "phoenix/client"
    "phoenix/ui"
    "bufio"
    "fmt"
    "os"
)


func main() {
    var client = client.NewClient(1)
    var hasUserStopped bool
    var scanner = bufio.NewScanner(os.Stdin)
    
    for !hasUserStopped {
        fmt.Println("Enter a command:")
        fmt.Printf(">: ")
        scanner.Scan()
        
        if scanner.Text() == "exit" {
            hasUserStopped = true
        } else {
            ui.ProcessUserCommands(client, scanner.Text()) 
        }
    }
    
    fmt.Println("Waiting for remaining processes to finish...")
    balancer.GetLoadBalancer().SyncGroup.Wait()
}
