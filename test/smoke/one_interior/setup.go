package smoke

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/ginkgo"
)

var OneRouterSpec = v1alpha1.InterconnectSpec{
	DeploymentPlan: v1alpha1.DeploymentPlanType{
		Size:      1,
		Image:     framework.TestContext.QdrImage,
		Role:      "interior",
		Placement: "Any",
	},
}

// SetupTopology will deploy operator and interconnect using a predefined Spec
func SetupTopology(f *framework.Framework, ic *v1alpha1.Interconnect, err error) {

	f = test.NewFramework("one-interior")

	// After operator deployed and before running tests
	ginkgo.JustBeforeEach(func() {
		ic, err = test.CreateInterconnect(f.GetFirstContext(), 1, OneRouterSpec)
		// Verify deployment worked
		if err == nil {
			// Verify Interconnect is running
			err = framework.WaitForDeployment(f.GetFirstContext().Clients.KubeClient, f.GetFirstContext().Namespace, ic.Name, 1, framework.RetryInterval, framework.Timeout)
		}

		return
	})

	return

}
