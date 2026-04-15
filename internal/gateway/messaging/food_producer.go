package messaging

import (
	"golang-clean-architecture/internal/model"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type FoodProducer struct {
	Producer[*model.FoodEvent]
}

func NewFoodProducer(producer sarama.SyncProducer, log *logrus.Logger) *FoodProducer {
	return &FoodProducer{
		Producer: Producer[*model.FoodEvent]{
			Producer: producer,
			Topic:    "food",
			Log:      log,
		},
	}
}
