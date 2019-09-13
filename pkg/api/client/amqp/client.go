package amqp

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	v1 "k8s.io/api/core/v1"
)

type Client interface {
	Deploy(context framework.ContextData, pod v1.Pod)
	Status()
	Stop()
}

type (
	ClientStatus int
)
