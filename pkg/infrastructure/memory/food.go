package memory

import "github.com/brewerywiwis/food-review-backend/pkg/domain/food"

type FoodStorage struct {
	foodID int64
	foods  []food.Food
}

func (s *FoodStorage) GetAll() []food.Food {
	return s.foods
}

func (s *FoodStorage) GetByID(id string) food.Food {
	return food.Food{}
}
func (s *FoodStorage) GetByName(name string) food.Food {
	return food.Food{}
}
func (s *FoodStorage) Create(food food.Food) (food.Food, error) {
	food.ID = s.foodID
	s.foodID++
	s.foods = append(s.foods, food)
	return food, nil
}
func (s *FoodStorage) Update(food food.Food) (food.Food, error) {
	return food, nil
}
func (s *FoodStorage) Delete(food food.Food) error {
	return nil
}
