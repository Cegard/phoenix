package components


type Balancer interface {
    addService (*service)
    removeIdleServices()
    AssignRequest (*request) *Response
}
