package router_broker_test

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"testing"

	. "github.com/onsi/ginkgo"
)

func TestRouterBroker(t *testing.T) {
	test.Initialize()
	RunSpecs(t, "RouterBroker Suite")
}
