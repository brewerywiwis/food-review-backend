package memory

import (
	"time"

	"github.com/brewerywiwis/food-review-backend/pkg/domain/food"
)

// Food is alias name for food.Food
type Food = food.Food

// ID is alias name for food.ID
type ID = food.ID

// FoodStorage is represent the infrastructure of food storage
type FoodStorage struct {
	foods []Food
}

// GetAll is used to retrieve all food in the memory db
func (s *FoodStorage) GetAll() []Food {
	return s.foods
}

// GetByID is used to retrieve a food by ID in the memory db
func (s *FoodStorage) GetByID(id ID) Food {
	return Food{}
}

// GetByName is used to retrieve a food by name in the memory db
func (s *FoodStorage) GetByName(name string) Food {
	return Food{}
}

// Create is used to add a new food record to the memory db
func (s *FoodStorage) Create(food Food) (Food, error) {
	if len(s.foods) <= 0 {
		food.ID = 1
	} else {
		food.ID = s.foods[len(s.foods)-1].ID + 1
	}
	food.CreatedAt = time.Now()
	s.foods = append(s.foods, food)
	return food, nil
}

// Update is used to update a food record
func (s *FoodStorage) Update(food Food) (Food, error) {
	return food, nil
}

// Delete is used to delete a food record
func (s *FoodStorage) Delete(food Food) error {
	return nil
}
