package utils

const MIN_PROCESS_TIME = 5
const MAX_PROCESS_TIME = 10
const MIN_RUNING_SERVICES = 1
const MAX_SERVICE_CAPACITY = 10
const SUCCESS_PROBABILITY float64 = 0.99
const (
    WAITING_STATUS uint = iota
    SUCCEEDED_STATUS
    FAILED_STATUS    
)
