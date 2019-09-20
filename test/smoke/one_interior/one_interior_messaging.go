package one_interior

import (
	"fmt"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/api/client/amqp"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/api/client/amqp/qeclients/python"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"time"
)

var _ = ginkgo.Describe("OneInteriorMessaging", func() {
	// Clients to be initialized for each test
	var (
		ctx      *framework.ContextData
		sender   amqp.Client
		receiver amqp.Client
		url      string
		err      error
	)

	// Run the clients
	ginkgo.JustBeforeEach(func() {
		ctx = Framework.GetFirstContext()

		// Url to use on both clients
		url = fmt.Sprintf("amqp://%s:5672/one_interior_messaging", DeployName)

		// Building sender client
		senderBuilder := &python.AmqpPythonSenderBuilder{}
		senderBuilder.New("sender", *ctx, url)
		sender, err = senderBuilder.Timeout(60).Messages(100).MessageContent("ABCDEFG").Build()
		gomega.Expect(err).To(gomega.BeNil())

		// Building receiver client
		receiverBuilder := &python.AmqpPythonReceiverBuilder{}
		receiverBuilder.New("receiver", *ctx, url)
		receiver, err = receiverBuilder.Timeout(60).Messages(100).Build()
		gomega.Expect(err).To(gomega.BeNil())
	})

	ginkgo.It("Exchange messages through router mesh", func() {
		// Deploying clients
		err = sender.Deploy()
		gomega.Expect(err).To(gomega.BeNil())

		err = receiver.Deploy()
		gomega.Expect(err).To(gomega.BeNil())

		// TODO Add a "Wait" method to amqp.Client
		time.Sleep(30 * time.Second)

		// Validating results
		senderResult := sender.Result()
		receiverResult := receiver.Result()

		// Ensure results obtained
		gomega.Expect(senderResult).NotTo(gomega.BeNil())
		gomega.Expect(receiverResult).NotTo(gomega.BeNil())

		// Validate sent/received messages
		gomega.Expect(senderResult.Delivered).To(gomega.Equal(100))
		gomega.Expect(receiverResult.Delivered).To(gomega.Equal(100))
	})
})
