package entity

type Food struct {
	ID          string  `gorm:"column:id;primaryKey"`
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Ingredients string  `gorm:"column:ingredients"`
	Price       int     `gorm:"column:price"`
	Rate        float64 `gorm:"column:rate"`
	Types       string  `gorm:"column:types"`
	PicturePath string  `gorm:"column:picture_path"`
	CreatedAt   int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (f *Food) TableName() string {
	return "food"
}
