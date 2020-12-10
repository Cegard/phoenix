# phoenix
System networks simulation

Usage:

Run `go run main.go` commands:

 - `"send n"`: Sends a number `n` of requests to the load balancer, example: `send 1000` sends 1000 requests to the server.
 - `"balancerStatus"`: Prints the status of the server with the information of each of the running services.
 - `"serviceStatus id"`: Prints the information of the service with id `"id"`, example `serviceStatus 1` prints out the information of the service with the id 1.
 - `exit`: Exits from the program.