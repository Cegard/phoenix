package utils

import (
    "math/rand"
    "time"
)


func RandomInt(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return (rand.Intn(max - min) + min)
}


func RandomFloat() float64 {
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
}


func FindIndexBy (list []*interface{}, rule func(*interface{}) bool) (int, bool) {
    var index = 0
    var listLength = len(list)
    
    for index < listLength && !rule(list[index]) {
        index++
    }
    
    if index == listLength {
        
        return index, false
    }
    
    return index, true
}