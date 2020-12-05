package server

import (
    "sync"
)


type Load struct {
    sync.Mutex
    value int
}


func (load *Load) addToValue (toAdd int) {
    load.value = load.value + toAdd
}


func (load *Load) IncreaseLoad() {
    load.Lock()
    load.addToValue(1)
    load.Unlock()
}


func (load *Load) DecreaseLoad() {
    load.Lock()
    load.addToValue(-1)
    load.Unlock()
}


func (load *Load) GetValue() int {
    defer load.Unlock()
    load.Lock()
    
    return load.value
}
