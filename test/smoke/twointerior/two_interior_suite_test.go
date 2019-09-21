package twointerior_test

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"testing"

	. "github.com/onsi/ginkgo"
)

func TestTwoInterior(t *testing.T) {
	test.Initialize()
	RunSpecs(t, "TwoInterior Suite")
}
