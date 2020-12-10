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
        fmt.Printf("\nEnter a command:\n")
        fmt.Printf(">: ")
        scanner.Scan()
        
        if scanner.Text() == "exit" {
            hasUserStopped = true
        } else {
            var message, err = ui.ProcessUserCommands(client, scanner.Text())
            
            if err == nil {
                fmt.Printf("\n%s\n", message)
            } else {
                fmt.Println(fmt.Errorf("%w", err))
            }
        }
    }
    
    fmt.Printf("\nWaiting for remaining processes to finish...\n")
    balancer.GetLoadBalancer().SyncGroup.Wait()
}
