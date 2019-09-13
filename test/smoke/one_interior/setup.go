package one_interior

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

const (
	DeployName = "one-interior"
	DeploySize = 1
)

var (
	Framework *framework.Framework
	IcSpec    *v1alpha1.InterconnectSpec
)

// Create the Framework instance to be used one_interior tests
var _ = ginkgo.BeforeEach(func() {
	// Setup the topology
	Framework = framework.NewFramework("one-interior", framework.TestContext.GetContexts()[0])
}, 60)

// Deploy Interconnect
var _ = ginkgo.JustBeforeEach(func() {
	// Context to use
	ctx := Framework.GetFirstContext()

	// Deploy the Interconnect instance before running tests
	IcSpec = &v1alpha1.InterconnectSpec{
		DeploymentPlan: v1alpha1.DeploymentPlanType{
			Size:      DeploySize,
			Image:     framework.TestContext.QdrImage,
			Role:      "interior",
			Placement: "Any",
		},
	}

	// After operator deployed and before running tests
	_, err := ctx.CreateInterconnectFromSpec(1, DeployName, *IcSpec)
	gomega.Expect(err).To(gomega.BeNil())

	// Verify deployment worked
	// Verify Interconnect is running
	err = framework.WaitForDeployment(ctx.Clients.KubeClient, ctx.Namespace, DeployName, DeploySize, framework.RetryInterval, framework.Timeout)
	gomega.Expect(err).To(gomega.BeNil())
})

// After each test completes, run cleanup actions to save resources
var _ = ginkgo.AfterEach(func() {
	Framework.AfterEach()
})
