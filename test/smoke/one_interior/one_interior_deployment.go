package one_interior

import (
	"fmt"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("OneInteriorDeployment", func() {

	ginkgo.It("Validates deployment", func() {
		gomega.Expect(SmokeFramework).NotTo(gomega.BeNil())
		gomega.Expect(SmokeIc).NotTo(gomega.BeNil())
		gomega.Expect(SmokeFramework.GetFirstContext()).NotTo(gomega.BeNil())
		fmt.Printf("Namespace = %s\n", SmokeFramework.GetFirstContext().Namespace)
	})

	ginkgo.It("Validates deployed network", func() {
		//fmt.Printf("Namespace = %s\n", f.Namespace)
	})

	ginkgo.It("Ensure one interior router available", func() {
		//fmt.Printf("Namespace = %s\n", f.Namespace)
	})

})
