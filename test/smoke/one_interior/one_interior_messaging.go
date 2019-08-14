package smoke

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("OneInteriorMessaging", func() {
	f := test.NewFramework("one-interior")

	ginkgo.It("Exchange messages", func() {
		gomega.Expect(f).NotTo(gomega.BeNil())
	})
})
