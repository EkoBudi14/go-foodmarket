package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func TransactionToResponse(transaction *entity.Transaction) *model.TransactionResponse {
	return &model.TransactionResponse{
		ID:         transaction.ID,
		User:       UserToResponse(&transaction.User),
		Food:       FoodToResponse(&transaction.Food),
		Quantity:   transaction.Quantity,
		Total:      transaction.Total,
		Status:     transaction.Status,
		PaymentUrl: transaction.PaymentUrl,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}
}
