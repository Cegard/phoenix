package messages

import (
    "phoenix/utils"
    "time"
)


type Request struct {
    RespondTo func(*Response)
    TimeToProcess time.Duration
}


func NewRequest(respondTo func(*Response) ) *Request {
    
    return &Request {
        RespondTo: respondTo,
        TimeToProcess: time.Duration(utils.RandomInt(
            utils.MIN_PROCESS_TIME,
            utils.MAX_PROCESS_TIME),
        ),
    }
}
