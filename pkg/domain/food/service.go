package food

// Service represents for food service
type Service interface {
	AddFood(food ...Food) error
	ListAllFood() ([]Food, error)
}

// NewService creates a new service
func NewService(repo *Repository) Service {
	return &service{repo}
}

type service struct {
	repo *Repository
}

func (s *service) AddFood(food ...Food) error {
	return nil
}

func (s *service) ListAllFood() ([]Food, error) {
	return nil, nil
}
