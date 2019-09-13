package amqp

import "github.com/fgiorgetti/qpid-dispatch-go-tests/pkg/framework"

type Client interface {
	Deploy()
	Status() ClientStatus
	Running() bool
	Interrupt()
	Result() ResultData
}

// SenderBuilder minimalist sample builder for AMQP Senders
type SenderBuider interface {
	New(name string) SenderBuider
	Context(data framework.ContextData) SenderBuider
	Url(url string) SenderBuider
	Messages(count int) SenderBuider
	Timeout(timeout int) SenderBuider
	Param(name string, value string) SenderBuider
	MessageContent(content string) SenderBuider
	Build() (Client, error)
}

// ReceiverBuilder minimalist sample builder for AMQP Receivers
type ReceiverBuilder interface {
	New(name string) ReceiverBuilder
	Context(data framework.ContextData) ReceiverBuilder
	Url(url string) ReceiverBuilder
	Messages(count int) ReceiverBuilder
	Timeout(timeout int) ReceiverBuilder
	Param(name string, value string) ReceiverBuilder
	Build() (Client, error)
}

type Message struct {
	Address       string
	Content       string
	ContentSHA1   string
	Id            string
	CorrelationId string
	ReplyTo       string
	Expiration    int
	Priority      int
	Ttl           int
	UserId        string
}

//{'address': None, 'annotations': None, 'content': '1234567890', 'content_encoding': 'None', 'content_type': 'text/plain', 'correlation_id': None, 'creation_time': 0.000000,
// 'delivery_count': 0, 'durable': False, 'expiration': 0, 'first_acquirer': False, 'group_id': None, 'group_sequence': 0, 'id': None, 'inferred': False, 'instructions': None,
// 'priority': 4, 'properties': {}, 'reply_to': None, 'reply_to_group_id': None, 'subject': None, 'ttl': 0, 'user_id': ''}

type ResultData struct {
	Messages  []Message
	Delivered int
	Released  int
	Modified  int
}

type ClientStatus int

const (
	Starting ClientStatus = iota
	Running
	Success
	Error
	Timeout
	Interrupted
)
