package info

import (
    "sync"
    "fmt"
)


type Info struct {
    sync.Mutex
    entries map[int] string
}


func NewInfo() *Info {
    
    return &Info {
        entries: make(map[int] string),
    }
}


func (info *Info) RegisterStat (dataId int, data string) {
    info.Lock()
    info.entries[dataId] = data
    info.Unlock()
}


func  (info *Info) GetEntryInfo (dataId int) string {
    info.Lock()
    defer info.Unlock()
    
    return info.entries[dataId]
}


func (info *Info) GetAllInfo () []string {
    info.Lock()
    defer info.Unlock()
    
    var sliceIndex = 0;
    var arraySize = len(info.entries)
    var statuses = make([]string, arraySize)
    
    for mapIndex := range info.entries {
        statuses[sliceIndex] = fmt.Sprintf("%s", info.entries[mapIndex])
        sliceIndex++
    }
    
    return statuses
}
