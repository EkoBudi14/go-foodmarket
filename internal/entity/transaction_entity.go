package entity

type Transaction struct {
	ID         string `gorm:"column:id;primaryKey"`
	UserId     string `gorm:"column:user_id"`
	User       User   `gorm:"foreignKey:UserId;references:ID"`
	FoodId     string `gorm:"column:food_id"`
	Food       Food   `gorm:"foreignKey:FoodId;references:ID"`
	Quantity   int    `gorm:"column:quantity"`
	Total      int    `gorm:"column:total"`
	Status     string `gorm:"column:status"`
	PaymentUrl string `gorm:"column:payment_url"`
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (c *Transaction) TableName() string {
	return "transaction"
}
