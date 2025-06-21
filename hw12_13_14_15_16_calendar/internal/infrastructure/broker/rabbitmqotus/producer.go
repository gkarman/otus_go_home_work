package rabbitmqotus

import (
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/broker"
	"github.com/gkarman/otus_go_home_work/hw12_13_14_15_calendar/internal/infrastructure/config"
)

type rabbitProducer struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	cfg     *config.BrokerConf
}

func NewRabbitProducer(_ *config.BrokerConf) (broker.Producer, error) {
	return nil, nil
}
