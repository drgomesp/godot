package mvc

import (
	"net/http"

	"github.com/drgomesp/cargo/container"
)

// Action defines an interface for request/response handlers
type Action interface {
	Handler(http.ResponseWriter, *http.Request)
}

// ContainerAwareAction defines an action aware of the service container
type ContainerAwareAction struct {
	container *container.Container
}
