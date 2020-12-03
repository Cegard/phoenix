package messages

import (
    "phoenix/utils"
)


type Response struct {
    ServiceId int
    StatusCode uint
    Message string
}


func NewResponse (code uint, serviceId int) *Response {
    var message string
    
    if code == utils.SUCCEEDED_STATUS {
        message = "Ok."
    } else {
        message = "Oops! Something went wrong."
    }
    
    return &Response {
        ServiceId: serviceId,
        StatusCode: code,
        Message: message,
    }
}
