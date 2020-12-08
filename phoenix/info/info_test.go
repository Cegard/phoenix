package info

import (
    "testing"
    "github.com/stretchr/testify/assert"
)


func TestRegisterStat (t *testing.T) {
    var info = NewInfo()
    
    info.RegisterStat(0, "dummy stat")
    
    assert.NotNil(
        t,
        info.GetEntryInfo(0),
        "Info is not registering data",
    )
}


func TestGetEntryInfo (t *testing.T) {
    var info = NewInfo()
    var dummyData = "This is a dummy data"
    
    info.RegisterStat(0, dummyData)
    
    assert.Equal(
        t,
        dummyData,
        info.GetEntryInfo(0),
        "Info is not retrieving data",
    )
}


func TestGetAllInfo (t *testing.T) {
    var info = NewInfo()
    var dummyData = "This is a dummy data"
    var entriesNumber = 10
    
    for i := 0; i < entriesNumber; i++ {
        info.RegisterStat(i, dummyData)
    }
    
    assert.Equal(
        t,
        dummyData,
        info.GetAllInfo(dummyData)[0],
        "Info is not retrieving all data",
    )
    
    assert.Equal(
        t,
        entriesNumber + 1,
        len(info.GetAllInfo(dummyData)),
        "Info is not retrieving all data",
    )
}