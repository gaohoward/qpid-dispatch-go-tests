package router_broker_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRouterBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RouterBroker Suite")
}
