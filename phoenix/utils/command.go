package utils

type Command int

const (
    Send Command = iota
    ServerStatus
    ServiceStatus
)


func (c Command) String() string{
    
    return [...]string{"send", "balancerStatus", "serviceStatus"}[c]
}
