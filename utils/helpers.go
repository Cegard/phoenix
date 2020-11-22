package utils

import (
    "math/rand"
    "time"
)


func RandomInt(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return (rand.Intn(max - min) + min)
}