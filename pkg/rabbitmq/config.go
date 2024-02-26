package rabbitmq

import (
	"ProducerAPI/src/configs"
	"ProducerAPI/src/pkg/utils"
)

type MQConfig struct {
	// Connection fields
	Hostname    string
	Port        int
	Username    string
	Password    string
	VirtualHost string

	// Routing Configuration
	DeadLetterExchangeName string
	DeadLetterQueueSuffix  string

	// Naming Configuration
	QueueNamingStrategy utils.QueueNamingStrategy
}

const DefaultDeadLetterExchangeName = "DLX"
const DefaultDeadLetterQueueSuffix = "dlq"

func NewConfig(
) *MQConfig {
	rabbitConfig := configs.GetConfigs().RabbitConfig
	return &MQConfig{
		Hostname:               rabbitConfig.Hostname,
		Port:                   rabbitConfig.Port,
		Username:               rabbitConfig.Username,
		Password:               rabbitConfig.Password,
		VirtualHost:            rabbitConfig.VirtualHost,
		DeadLetterExchangeName: DefaultDeadLetterExchangeName,
		DeadLetterQueueSuffix:  DefaultDeadLetterQueueSuffix,
		QueueNamingStrategy:    utils.DefaultQueueNamingStrategy{},
	}
}