package logic

import (
	"context"
	"encoding/json"
	"frozen-go-project/common/enum"
	"frozen-go-project/common/errors/business_errors"
	"frozen-go-project/common/mq_msg"
	"frozen-go-project/rpc/event-rpc/internal/svc"
	event_rpc "frozen-go-project/rpc/event-rpc/pb"
	"github.com/Shopify/sarama"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserActionLogic {
	return &UserActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserActionLogic) UserAction(in *event_rpc.UserActionReq) (*event_rpc.CommonResponse, error) {
	topic := in.Common.Topic
	if len(topic) <= 0 {
		return nil, business_errors.NoMqTopic
	}
	userActionMsg := &mq_msg.UserActionMsg{
		UserId:      in.UserAction.UserId,
		UserAction:  enum.UserAction(in.UserAction.UserAction),
		EventTimeMs: in.Common.EventTimeMs,
	}
	msg, _ := json.Marshal(userActionMsg)
	partition, offset, err := l.svcCtx.KafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	})
	if err != nil {
		return nil, err
	}
	return &event_rpc.CommonResponse{
		Partition: partition,
		Offset:    offset,
	}, nil
}
