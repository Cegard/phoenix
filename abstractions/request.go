package abstractions

import (
    _"fmt"
    "phoenix/utils"
    "time"
)


type request struct {
    Id int
    TimeToProcess time.Duration
}


func CreateRequest(id int) *request {
    
    return &request {
        Id: id,
        TimeToProcess: time.Duration(utils.RandomInt(
            utils.MIN_PROCESS_TIME,
            utils.MAX_PROCESS_TIME),
        ),
    }
}