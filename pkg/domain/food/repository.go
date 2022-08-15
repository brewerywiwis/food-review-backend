package food

// Repository is defined to be an adapter for infrastructure
type Repository interface {
	GetAll() []Food
	GetByID(id string) Food
	GetByName(name string) Food
	Create(food Food) (Food, error)
	Update(food Food) (Food, error)
	Delete(food Food) error
}
