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
