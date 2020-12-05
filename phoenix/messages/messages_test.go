package messages

import (
    "phoenix/utils"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestRequestCreation (t *testing.T) {
    var request = NewRequest(func (*Response){})
    
    assert.NotEqual(t, request.TimeToProcess, 0, "Request is not being created")
}


func TestResponseCreation (t *testing.T) {
    var response = NewResponse(utils.SucceededStatus, 0)
    
    assert.NotEqual(t, response.Message, "", "Response is not being created")
}
