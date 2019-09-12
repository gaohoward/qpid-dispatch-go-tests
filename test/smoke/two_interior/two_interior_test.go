package two_interior

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/qdrmanagement"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

/**
Validates the formed mesh
*/
var _ = Describe("Validates the formed mesh", func() {

	var (
		ctx1 *framework.ContextData
		ctx2 *framework.ContextData
	)

	// Initialize after frameworks have been created
	JustBeforeEach(func() {
		ctx1 = FrameworkQdrOne.GetFirstContext()
		ctx2 = FrameworkQdrOne.GetFirstContext()
	})

	It("Query routers in the network on each pod", func() {
		ValidateRoutersInNetwork(ctx1, QdrOneName, 2)
		ValidateRoutersInNetwork(ctx2, QdrTwoName, 2)

	})
})

// ValidateRoutersInNetwork uses qdmanage query to retrieve nodes in the router network.
// It iterates through all pods available in the provided context and deployment and waits
// till the expected amount of nodes are present or till it times out.
func ValidateRoutersInNetwork(ctx *framework.ContextData, deploymentName string, expectedCount int) {
	fmt.Printf("\n\nREACHED POINT 1\n\n")
	// Retrieves pods
	podList, err := ctx.ListPodsForDeploymentName(deploymentName)
	gomega.Expect(err).To(gomega.BeNil())
	gomega.Expect(podList.Items).To(gomega.HaveLen(1))
	fmt.Printf("\n\nREACHED POINT 2\n\n")

	// Iterate through pods and execute qdmanage query across all pods
	for _, pod := range podList.Items {
		fmt.Printf("\n\nREACHED POINT 3...\n\n")
		// Wait till expected amount of nodes are present or till it times out
		err := qdrmanagement.WaitForQdrNodesInPod(*ctx, pod, expectedCount, framework.RetryInterval, framework.Timeout)
		gomega.Expect(err).To(gomega.BeNil())
	}
}
