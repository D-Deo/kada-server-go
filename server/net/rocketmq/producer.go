package rocketmq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type SendResult = primitive.SendResult

type Producer struct {
	p rocketmq.Producer

	groupId     string
	nameSrv     string
	credentials *Credentials

	topic string
	tag   string
	keys  []string
}

func NewProducer(groupId string, nameSrv string) *Producer {
	p := &Producer{
		groupId: groupId,
		nameSrv: nameSrv,
	}
	return p
}

func (o *Producer) Use(topic string, tag string, keys []string) *Producer {
	o.topic = topic
	o.tag = tag
	o.keys = keys
	return o
}

func (o *Producer) UseCredentials(accessKey string, secretKey string) {
	if o.credentials == nil {
		o.credentials = &Credentials{}
	}
	o.credentials.AccessKey = accessKey
	o.credentials.SecretKey = secretKey
}

func (o *Producer) Start() error {
	var opts []producer.Option
	opts = append(opts,
		producer.WithGroupName(o.groupId),
		producer.WithNsResovler(primitive.NewPassthroughResolver([]string{o.nameSrv})),
		producer.WithRetry(2),
	)

	if o.credentials != nil {
		opts = append(opts, producer.WithCredentials(primitive.Credentials{
			AccessKey: o.credentials.AccessKey,
			SecretKey: o.credentials.SecretKey,
		}))
	}

	p, err := rocketmq.NewProducer(opts...)
	if err != nil {
		return err
	}

	if err := p.Start(); err != nil {
		return err
	}
	o.p = p

	return nil
}

func (o *Producer) SendMessage(data []byte) (*SendResult, error) {
	msg := primitive.NewMessage(o.topic, data)
	msg.WithTag(o.tag)
	msg.WithKeys(o.keys)

	ret, err := o.p.SendSync(context.Background(), msg)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
