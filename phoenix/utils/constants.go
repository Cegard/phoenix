package utils

const MinProcessTime = 3 // seconds
const MaxProcessTime = 6 //seconds
const MinRunningServices = 1
const MaxServiceCapacity = 10
const CouldNotProcessMsg = "Couldn't process"
const NotRecognizedMsg = "Command not recognized"
const SuccessProbability float64 = 0.99
const (
    SucceededStatus uint = iota
    FailedStatus  
)
