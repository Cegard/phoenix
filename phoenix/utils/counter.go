package utils

import (
    "sync"
)


type Counter struct {
    sync.Mutex
    count int
}


func (counter *Counter) addToCount (toAdd int) {
    counter.count = counter.count + toAdd
}


func (counter *Counter) IncreaseCount() {
    counter.Lock()
    counter.addToCount(1)
    counter.Unlock()
}


func (counter *Counter) DecreaseCount() {
    counter.Lock()
    counter.addToCount(-1)
    counter.Unlock()
}


func (counter *Counter) GetCount() int {
    defer counter.Unlock()
    counter.Lock()
    
    return counter.count
}
