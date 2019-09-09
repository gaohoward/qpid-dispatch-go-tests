package smoke

import (
	"fmt"
	"testing"

	_ "github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	_ "github.com/fgiorgetti/qpid-dispatch-go-tests/test/smoke/one_interior"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSmoke(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Smoke Suite")
}

var _ = BeforeEach(func() {
	fmt.Println("Before each Smoke Test")
}, 60)

var _ = AfterEach(func() {
	fmt.Println("After each Smoke Test")
}, 60)
