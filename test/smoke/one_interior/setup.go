package smoke

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/gomega"
)

func CreateInterconnectSpec() *v1alpha1.InterconnectSpec {
	return &v1alpha1.InterconnectSpec{
		DeploymentPlan: v1alpha1.DeploymentPlanType{
			Size:      1,
			Image:     framework.TestContext.QdrImage,
			Role:      "interior",
			Placement: "Any",
		},
	}
}
// SetupTopology will deploy interconnect using a predefined Spec
func SetupTopology(f *framework.Framework, interconnect *v1alpha1.InterconnectSpec) {

	// After operator deployed and before running tests
	ic, err := test.CreateInterconnect(f.GetFirstContext(), 1, *interconnect)
	gomega.Expect(err).To(gomega.BeNil())

	// Verify deployment worked
	// Verify Interconnect is running
	err = framework.WaitForDeployment(f.GetFirstContext().Clients.KubeClient, f.GetFirstContext().Namespace, ic.Name, 1, framework.RetryInterval, framework.Timeout)
	gomega.Expect(err).To(gomega.BeNil())

}
