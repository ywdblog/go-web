package dao

import (
	"context"

	"github.com/Shopify/sarama"
	skykafka "github.com/WiFeng/go-sky/kafka"
	"github.com/WiFeng/go-sky/log"
)

type Kafka struct {
}

func (k Kafka) KafkaPush(ctx context.Context, name string, topic string, value string) (bool, error) {
	p, err := skykafka.NewSyncProducer(ctx, name)

	if err != nil {
		log.Errorw(ctx, "skykafka.NewSyncProducer error", "err", err)
		return false, err
	}
	defer p.Close()

	msg := &sarama.ProducerMessage{
		// Key:   sarama.StringEncoder(strKey),
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	part, offset, err := p.SendMessageContext(ctx, msg)
	if err != nil {
		log.Errorw(ctx, "SendMessage error", "err", err)
	} else {
		log.Infow(ctx, "SendMessage ok", "part", part, "offset", offset)
	}
	return true, err
}
