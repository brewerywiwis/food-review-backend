package food

// Repository is defined to be an adapter for infrastructure
type Repository interface {
	GetByID(id string)
	GetByName(name string)
	Create(food Food) (Food, error)
	Update(food Food) (Food, error)
	Delete(food Food) error
}
