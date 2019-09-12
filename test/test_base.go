package test

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/framework/ginkgowrapper"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// init Initializes Ginkgo and parse command line arguments
func init() {
	fmt.Printf("\n\n\n\n\n\nINIT HAS BEEN CALLED\n\n\n\n\n")
	framework.HandleFlags()
	gomega.RegisterFailHandler(ginkgowrapper.Fail)
}

// Before suite common setup (happens only once per test suite)
var _ = ginkgo.SynchronizedBeforeSuite(func() []byte {
	// Unique initialization (node 1 only)
	fmt.Println("Base test setup - only happens once per test suite")
	return nil
}, func(data []byte) {
	// Initilization for each parallel node
}, 10)

// After suite common teardown (happens only once per test suite)
var _ = ginkgo.SynchronizedAfterSuite(func() {
	// All nodes tear down
}, func() {
	// Node1 only tear down
	fmt.Println("Base test teardown - only happens once per test suite")
}, 10)
