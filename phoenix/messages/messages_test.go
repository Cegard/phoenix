package messages

import (
    "testing"
    "phoenix/utils"
)


func TestRequestCreation (t *testing.T) {
    var request = NewRequest(func (*Response){})
    
    if request.TimeToProcess == 0 {
        t.Errorf("Request is not being created")
    }
}


func TestResponseCreation (t *testing.T) {
    var response = NewResponse(utils.SucceededStatus, 0)
    
    if response.Message == "" {
        t.Errorf("Response is not being created")
    }
}
