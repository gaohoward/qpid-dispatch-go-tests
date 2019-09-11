package one_interior_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOneInterior(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OneInterior Suite")
}
