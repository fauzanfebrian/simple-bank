package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	Client *asynq.Client
}

func NewRedisTaskDistributor(option asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(option)
	return &RedisTaskDistributor{
		Client: client,
	}
}
