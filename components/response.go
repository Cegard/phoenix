package components

import (
    "phoenix/utils"
)


type Response struct {
    StatusCode uint
    Message string
}


func CreateResponse (code uint) *Response {
    var message string
    
    if code == utils.SUCCEEDED_STATUS {
        message = "Ok."
    } else {
        message = "Oops! Something went wrong."
    }
    
    return &Response {
        StatusCode: code,
        Message: message,
    }
}
