package main

import (
    "phoenix/components"
    "fmt"
    "os"
    "bufio"
)


func main() {
    var client = components.CreateClient(1)
    var hasUserStopped bool
    var in = bufio.NewScanner(os.Stdin)
    
    for !hasUserStopped {
        fmt.Println("Enter a command:")
        fmt.Printf(">: ")
        in.Scan()
        
        if in.Text() == "exit" {
            hasUserStopped = true
        } else {
            components.ProcessUserCommands(client, in.Text()) 
        }
    }
    
    fmt.Println("Waiting for remaining processes to finish...")
    components.GetLoadBalancer().SyncGroup.Wait()
}
