package utils

import (
    "sync"
)


type Counter struct {
    sync.Mutex
    count int
}


func (counter *Counter) addToCount (toAdd int) {
    counter.Lock()
    counter.count = counter.count + toAdd
    counter.Unlock()
}


func (counter *Counter) IncreaseCount() {
    counter.addToCount(1)
}


func (counter *Counter) DecreaseCount() {
    counter.addToCount(-1)
}


func (counter *Counter) GetCount() int {
    defer counter.Unlock()
    counter.Lock()
    
    return counter.count
}
