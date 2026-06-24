package repository

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type FoodRepository struct {
	Repository[entity.Food]
	Log *logrus.Logger
}

func NewFoodRepository(log *logrus.Logger) *FoodRepository {
	return &FoodRepository{
		Log: log,
	}
}

func (r *FoodRepository) FindByid(db *gorm.DB, food *entity.Food, id string) error {
	return db.Where("id = ? ", id).Take(food).Error
}

func (r *FoodRepository) Search(db *gorm.DB, request *model.FoodRequest) ([]entity.Food, error) {
	var foods []entity.Food

	if err := db.Scopes(r.FilterFood(request)).Limit(request.Limit).Find(&foods).Error; err != nil {
		return nil, err
	}

	return foods, nil
}

func (r *FoodRepository) FilterFood(request *model.FoodRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {

		// Filter by ID (find single food)
		if request.ID != "" {
			tx = tx.Where("id = ?", request.ID)
		}

		// Filter by Name (LIKE)
		if request.Name != "" {
			tx = tx.Where("name LIKE ?", "%"+request.Name+"%")
		}

		// Filter by Types (LIKE)
		if request.Types != "" {
			tx = tx.Where("types LIKE ?", "%"+request.Types+"%")
		}

		// Filter by Price range
		if request.PriceFrom > 0 {
			tx = tx.Where("price >= ?", request.PriceFrom)
		}
		if request.PriceTo > 0 {
			tx = tx.Where("price <= ?", request.PriceTo)
		}

		// Filter by Rate range
		if request.RateFrom > 0 {
			tx = tx.Where("rate >= ?", request.RateFrom)
		}
		if request.RateTo > 0 {
			tx = tx.Where("rate <= ?", request.RateTo)
		}

		return tx
	}
}
