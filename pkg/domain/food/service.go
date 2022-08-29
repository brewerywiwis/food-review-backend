package food

// Service represents for food service
type Service interface {
	ListAllFood() []Food
	AddFood(food ...Food) error
	UpdateFood(food Food) error
	DeleteFood(id ID) error
}

// NewService creates a new service
func NewService(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repository Repository
}

func (s *service) ListAllFood() []Food {
	return s.repository.GetAll()
}

func (s *service) AddFood(food ...Food) error {
	for _, f := range food {
		if _, err := s.repository.Create(f); err != nil {
			return err
		}
	}
	return nil
}

func (s *service) UpdateFood(food Food) error {
	if food.ID == nil {

	}
	return nil
}
func (s *service) DeleteFood(id ID) error {
	return nil
}
