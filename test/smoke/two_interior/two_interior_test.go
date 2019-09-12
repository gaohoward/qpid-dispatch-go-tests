package two_interior

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/qdrmanagement"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("TwoInterior", func() {

	var (
		ctx1 *framework.ContextData
		ctx2 *framework.ContextData
	)

	// Initialize after frameworks have been created
	JustBeforeEach(func() {
		ctx1 = FrameworkQdrOne.GetFirstContext()
		ctx2 = FrameworkQdrOne.GetFirstContext()
	})

	It("Deploys two interior routers in different clusters", func() {
		gomega.Expect(FrameworkQdrOne).NotTo(gomega.BeNil())
		gomega.Expect(FrameworkQdrTwo).NotTo(gomega.BeNil())

		// Verify deployment present (create a method for that)
		depList, err := ctx1.Clients.KubeClient.Apps().Deployments(ctx1.Namespace).List(v1.ListOptions{})
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(depList).NotTo(gomega.BeNil())

		depList, err = ctx2.Clients.KubeClient.Apps().Deployments(ctx2.Namespace).List(v1.ListOptions{})
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(depList).NotTo(gomega.BeNil())
	})

	FIt("Query routers in the network on each pod", func() {
		depQdrOne, err := ctx1.GetDeployment(QdrOneName)
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(depQdrOne).NotTo(gomega.BeNil())

		podList, err := ctx1.ListPodsForDeployment(depQdrOne)
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(podList.Items).To(gomega.HaveLen(1))
		for _, pod := range podList.Items {
			err := qdrmanagement.WaitForQdrNodesInPod(*ctx1, pod, 2, framework.RetryInterval, framework.Timeout)
			gomega.Expect(err).To(gomega.BeNil())

			nodes, err := qdrmanagement.QdmanageQueryNodes(*ctx1, pod.Name)
 			//connections, err := qdrmanagement.ListInterRouterConnectionsForPod(*ctx1, pod)
			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(nodes).To(gomega.HaveLen(2))
			//gomega.Expect(len(connections)).To(gomega.Equal(2))
		}

	})
})
