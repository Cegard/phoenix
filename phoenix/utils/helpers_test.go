package utils

import (
    "fmt"
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestGetCount (t *testing.T) {
    var counter = &Counter{}
    
    assert.Equal(
        t,
        0,
        counter.GetCount(),
        "Not getting th current count",
    )
}


func TestIncreaseCount (t *testing.T) {
    var counter = &Counter{}
    
    oldCount := counter.GetCount()
    counter.IncreaseCount()
    newCount := counter.GetCount()
    
    assert.Greater(
        t,
        newCount,
        oldCount,
        "Counter's not increasing count",
    )
}


func TestDecreaseCount (t *testing.T) {
    var counter = &Counter{}
    
    oldCount := counter.GetCount()
    counter.IncreaseCount()
    counter.DecreaseCount()
    
    assert.Equal(
        t,
        oldCount,
        counter.GetCount(),
        "Counter's not decreasing count",
    )
}


func TestRandomInt (t *testing.T) {
    var randomInt = RandomInt(MinProcessTime, MaxProcessTime)
    
    assert.GreaterOrEqual(
        t,
        randomInt,
        MinProcessTime,
        "Random int not in boundaries",
    )
    assert.LessOrEqual(
        t,
        randomInt,
        MaxProcessTime,
        "Random int not in boundaries",
    )
}


func TestRandomFloat (t *testing.T) {
    assert.Equal(
        t,
        fmt.Sprintf("%T", float64(0.0)),
        fmt.Sprintf("%T", RandomFloat()),
        "Random not float64",
    )
}


func TestNotNumberError (t *testing.T) {
    var notNumberError = &NotNumberError{originalValue: "#"}
    
    assert.Error(
        t,
        notNumberError,
        "It doesn't implements Error",
    )
    
    assert.NotEmpty(
        t,
        fmt.Errorf("%w", notNumberError),
        "It doesn't show get the error",
    )
}
