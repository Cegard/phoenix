package messages

import (
    "phoenix/utils"
    "time"
)


type Request struct {
    RespondTo func(*Response)
    TimeToProcess time.Duration
}


func NewRequest(respondTo func (*Response) ) *Request {
    
    return &Request {
        RespondTo: respondTo,
        TimeToProcess: time.Duration(utils.RandomInt(
            utils.MinProcessTime,
            utils.MaxProcessTime),
        ),
    }
}
