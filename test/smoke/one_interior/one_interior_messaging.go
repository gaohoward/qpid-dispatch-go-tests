package one_interior

import (
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("OneInteriorMessaging", func() {
	ginkgo.It("Exchange messages", func() {
		gomega.Expect(Framework).NotTo(gomega.BeNil())
	})
})
