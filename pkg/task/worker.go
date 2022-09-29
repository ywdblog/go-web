package task

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	skykafka "github.com/WiFeng/go-sky/kafka"
	"github.com/WiFeng/go-sky/log"
	"github.com/xiwujie/article/pkg/config"
	"github.com/xiwujie/article/pkg/entity"
)

const (
	articleKafkaSyncName = "article_kafka_sync"
)

// ============================
// 启动文章版权logo任务
// ============================

func startSyncWorker() {
	ctx := context.Background()
	log.Infow(ctx, "start consume kafka")
	group, err := skykafka.NewConsumerGroup(ctx, articleKafkaSyncName)
	if err != nil {
		log.Fatalf(ctx, "create consumer group error. err:%s", err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			log.Errorf(ctx, "group consume error. err:%s", err)
		}
	}()

	// Iterate over consumer sessions.
	for {
		topics := config.GlobalAppConfig.App.ConsumerSyncTopics
		handler := consumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			log.Errorf(ctx, "group consume error. err:%s", err)
			time.Sleep(30 * time.Second)
		}
	}
}

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := context.Background()
	for msg := range claim.Messages() {
		sess.MarkMessage(msg, "")
		log.Infow(ctx, "topic info", "topic", msg.Topic, "partition", msg.Partition, "offset", msg.Offset,
			"key", string(msg.Key), "value", string(msg.Value))
		var flatMessage entity.FlatMessage
		if err := json.Unmarshal(msg.Value, &flatMessage); err != nil {
			log.Errorw(ctx, "json Unmarshal error:"+string(msg.Value))
			continue
		}

		// 检查更新字段，如果更新字段不在审核范围内，则直接忽略。
		// 删除操作直接忽略；更新、插入则在各自的内部进行检测
		if flatMessage.IsDel || flatMessage.Type == entity.BinlogTypeDelete {
			continue
		}

	}
	return nil
}
