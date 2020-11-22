package abstractions

import (
    "time"
)


type Service struct {
    id int
}


func (service Service) processRequest (clientRequest request) {
    time.Sleep(time.Second * clientRequest.TimeToProcess)
}