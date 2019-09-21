package test

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework/ginkgowrapper"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// Initialize Ginkgo and parse command line arguments
// Ginkgo needs flag parser to be defined in the init() method
// which only gets called once.
// Although it is not doing anything, by calling it, it causes the "init()" to be invoked by Ginkgo,
// So it is important to call it.
func Initialize() {
	// This method here can be used to perform other common initialization for other suites.
}

func init() {
	framework.HandleFlags()
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
}

// Before suite validation setup (happens only once per test suite)
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Unique initialization (node 1 only)
	return nil
}, func(data []byte) {
	// Initialization for each parallel node
}, 10)

// After suite validation teardown (happens only once per test suite)
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// All nodes tear down
}, func() {
	// Node1 only tear down
	framework.RunCleanupActions(framework.AfterEach)
	framework.RunCleanupActions(framework.AfterSuite)
}, 10)
