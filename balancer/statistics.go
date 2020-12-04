package balancer

import (
    "sync"
)


type Info struct {
    sync.Mutex
    Entries map[int] string
}


func NewInfo() *Info {
    
    return &Info {
        Entries: make(map[int] string),
    }
}


func (info *Info) RegisterStat (dataId int, data string) {
    info.Lock()
    info.Entries[dataId] = data
    info.Unlock()
}
