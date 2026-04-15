package messaging

import (
	"golang-clean-architecture/internal/model"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type TransactionProducer struct {
	Producer[*model.TransactionEvent]
}

func NewTransactionProducer(producer sarama.SyncProducer, log *logrus.Logger) *TransactionProducer {
	return &TransactionProducer{
		Producer: Producer[*model.TransactionEvent]{
			Producer: producer,
			Topic:    "transactions",
			Log:      log,
		},
	}
}
