package oneinterior_test

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/onsi/ginkgo"
	"testing"
)

func TestOneInterior(t *testing.T) {
	test.Initialize()
	ginkgo.RunSpecs(t, "OneInterior Suite")
}
