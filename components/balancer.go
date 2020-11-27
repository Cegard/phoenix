package components


type Balancer interface {
    AddService (*service)
    RemoveService (int)
    AssignRequest (*request) *Response
}
