package one_interior

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/validation/qpiddispatch/management"
	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("OneInteriorDeployment", func() {

	var (
		ctx1 *framework.ContextData
	)

	// Initialize after framework has been created
	ginkgo.JustBeforeEach(func() {
		ctx1 = Framework.GetFirstContext()
	})

	ginkgo.It("Query routers in the network on each pod", func() {
		management.ValidateRoutersInNetwork(ctx1, DeployName, DeploySize)
	})

})
