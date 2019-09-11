package two_interior_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTwoInterior(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TwoInterior Suite")
}
