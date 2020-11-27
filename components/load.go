package components

import (
    "sync"
)


type Load struct {
    sync.Mutex
    Value int
}


func (load *Load) addToValue (toAdd int) {
    load.Value = load.Value + toAdd
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
