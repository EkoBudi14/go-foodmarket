package usecase

import (
	"golang-clean-architecture/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	TransactionRepository *repository.TransactionRepository
}

func NewTransactionUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate,
	transactionRepository *repository.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		DB:                    db,
		Log:                   log,
		Validate:              validate,
		TransactionRepository: transactionRepository,
	}
}
