package one_interior

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/test"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

const DeployName = "one-interior"

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
	ic, err := test.CreateInterconnect(f.GetFirstContext(), 1, DeployName, *interconnect)
	gomega.Expect(err).To(gomega.BeNil())

	// Verify deployment worked
	// Verify Interconnect is running
	err = framework.WaitForDeployment(f.GetFirstContext().Clients.KubeClient, f.GetFirstContext().Namespace, ic.Name, 1, framework.RetryInterval, framework.Timeout)
	gomega.Expect(err).To(gomega.BeNil())

}

var (
	SmokeFramework *framework.Framework
	SmokeIc *v1alpha1.InterconnectSpec
)

// Create the Framework instance to be used one_interior tests
var _ = ginkgo.BeforeEach(func() {
	fmt.Println("Creating framework for one-interior testing")

	// Setup the topology
	SmokeFramework = framework.NewFramework("one-interior", framework.TestContext.GetContexts()[0])
}, 60)


// Deploy Interconnect
var _ = ginkgo.JustBeforeEach(func() {
	// Deploy the Interconnect instance before running tests
	fmt.Println("Deploying InterconnectSpec for one-interior")
	SmokeIc = CreateInterconnectSpec()
	SetupTopology(SmokeFramework, SmokeIc)
})

// Clean up after once each one_interior test is done
var _ = ginkgo.AfterEach(func() {
	fmt.Println("Cleaning up framework for one-interior testing")
	// Clean up
	SmokeFramework.AfterEach()
}, 60)
