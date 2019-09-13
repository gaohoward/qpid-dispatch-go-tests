package python

import (
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/api/client/amqp"
	"github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"
)

// AmqpPythonSender amqp Client implementation
type AmqpPythonSender struct {
	context framework.ContextData
	name string
	url string
	messages int
	content string
	timeout int
	params []string
	values []string
}

func (a AmqpPythonSender) Deploy() {
	panic("implement me")
}

func (a AmqpPythonSender) Status() amqp.ClientStatus {
	panic("implement me")
}

func (a AmqpPythonSender) Running() bool {
	panic("implement me")
}

func (a AmqpPythonSender) Interrupt() {
	panic("implement me")
}

func (a AmqpPythonSender) Result() amqp.ResultData {
	panic("implement me")
}

// AmqpPythonSenderBuilder amqp SenderBuilder implementation
type AmqpPythonSenderBuilder struct {
	pythonSender AmqpPythonSender
}

func (a AmqpPythonSenderBuilder) New(name string) amqp.SenderBuider {
	a.pythonSender = AmqpPythonSender{
		name: name,
	}
	return a
}

func (a AmqpPythonSenderBuilder) Context(data framework.ContextData) amqp.SenderBuider {
	a.pythonSender.context = data
	return a
}

func (a AmqpPythonSenderBuilder) Url(url string) amqp.SenderBuider {
	a.pythonSender.url = url
	return a
}

func (a AmqpPythonSenderBuilder) Messages(count int) amqp.SenderBuider {
	a.pythonSender.messages = count
	return a
}

func (a AmqpPythonSenderBuilder) Timeout(timeout int) amqp.SenderBuider {
	a.pythonSender.timeout = timeout
	return a
}

func (a AmqpPythonSenderBuilder) Param(name string, value string) amqp.SenderBuider {
	a.pythonSender.params = append(a.pythonSender.params)
	a.pythonSender.values = append(a.pythonSender.values)
	return a
}

func (a AmqpPythonSenderBuilder) MessageContent(content string) amqp.SenderBuider {
	a.pythonSender.content = content
	return a
}

func (a AmqpPythonSenderBuilder) Build() (amqp.Client, error) {
	//TODO Add some validation to ensure all required have been populated
	// and return an error instead of nil
	return a.pythonSender, nil
}
