package entity

// User is a struct that represents a user entity
type User struct {
	ID          string  `gorm:"column:id;primaryKey"`
	Password    string  `gorm:"column:password"`
	Name        string  `gorm:"column:name"`
	Email       string  `gorm:"column:email"`
	Roles       *string `gorm:"column:roles"`
	Address     *string `gorm:"column:addres"`
	HouseNumber *string `gorm:"column:house_number"`
	Token       string  `gorm:"column:token"`
	PhoneNumber *string `gorm:"column:phone_number"`
	CreatedAt   int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *User) TableName() string {
	return "users"
}
