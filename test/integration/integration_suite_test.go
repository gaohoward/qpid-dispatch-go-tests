package integration_test

import (
	"fmt"
	"testing"

	_ "github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	_ "github.com/fgiorgetti/qpid-dispatch-go-tests/test/integration/router_broker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeEach(func() {
	fmt.Println("Before each Integration test")
}, 10)

var _ = AfterEach(func() {
	fmt.Println("After each Integration test")
}, 10)
