package smoke

import (
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
}, 60)

var _ = AfterEach(func() {
	//fmt.Println("After each Smoke Test")
}, 60)
