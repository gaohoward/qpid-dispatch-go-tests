package smoke

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("OneInteriorDeployment", func() {
	f := test.NewFramework("one-interior")
	ginkgo.It("Deploys one interior router", func() {
		gomega.Expect(f).NotTo(gomega.BeNil())
	})
})
