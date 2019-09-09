package smoke

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/interconnectedcloud/qdr-operator/pkg/apis/interconnectedcloud/v1alpha1"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("OneInteriorDeployment", func() {

	var (
		f framework.Framework
		ic v1alpha1.Interconnect
		err error
	)

	SetupTopology(&f, &ic, err)

	ginkgo.It("Validates deployment", func() {
		gomega.Expect(err).NotTo(gomega.HaveOccurred())
		gomega.Expect(f).NotTo(gomega.BeNil())
		gomega.Expect(ic).NotTo(gomega.BeNil())
		gomega.Expect(f.GetFirstContext()).NotTo(gomega.BeNil())
		fmt.Printf("Namespace = %s\n", f.GetFirstContext().Namespace)
	})
	ginkgo.It("Validates deployed network", func() {
		//fmt.Printf("Namespace = %s\n", f.Namespace)
	})
	ginkgo.It("Ensure one interior router available", func() {
		//fmt.Printf("Namespace = %s\n", f.Namespace)
	})
})
