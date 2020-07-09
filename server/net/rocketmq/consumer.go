package rocketmq

import (
    "context"
    "github.com/apache/rocketmq-client-go/v2"
    "github.com/apache/rocketmq-client-go/v2/consumer"
    "github.com/apache/rocketmq-client-go/v2/primitive"
)

type ConsumeResult int

const (
    ConsumeSuccess ConsumeResult = iota
    ConsumeRetryLater
    Commit
    Rollback
    SuspendCurrentQueueAMoment
)

type MessageExt = primitive.MessageExt

type MessageHandler func(ctx context.Context, msgs ...*MessageExt) (ConsumeResult, error)

type Consumer struct {
    GroupId     string
    NameSrv     string
    Topic       string
    Credentials *Credentials
    Handler     MessageHandler
}

func NewConsumer(groupId string, nameSrv string, topic string) *Consumer {
    c := &Consumer{
        GroupId: groupId,
        NameSrv: nameSrv,
        Topic:   topic,
    }
    return c
}

func (o *Consumer) UseCredentials(accessKey string, secretKey string) {
    if o.Credentials == nil {
        o.Credentials = &Credentials{}
    }
    o.Credentials.AccessKey = accessKey
    o.Credentials.SecretKey = secretKey
}

func (o *Consumer) Start(handler MessageHandler) error {
    var opts []consumer.Option
    opts = append(opts, consumer.WithGroupName(o.GroupId),
        consumer.WithNsResovler(primitive.NewPassthroughResolver([]string{o.NameSrv})),
        consumer.WithConsumerModel(consumer.Clustering),
    )

    if o.Credentials != nil {
        opts = append(opts, consumer.WithCredentials(primitive.Credentials{
            AccessKey: o.Credentials.AccessKey,
            SecretKey: o.Credentials.SecretKey,
        }))
    }
    opts = append(opts, consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset))

    c, err := rocketmq.NewPushConsumer(opts...)
    if err != nil {
        return err
    }

    if err := c.Subscribe(o.Topic, consumer.MessageSelector{}, o.onMessage); err != nil {
        return err
    }

    o.Handler = handler
    if err := c.Start(); err != nil {
        return err
    }

    return nil
}

func (o *Consumer) onMessage(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
    if o.Handler == nil {
        return consumer.ConsumeRetryLater, nil
    }
    ret, err := o.Handler(ctx, msgs...)
    if err != nil {
        return consumer.ConsumeSuccess, err
    }
    return consumer.ConsumeResult(ret), nil
}
