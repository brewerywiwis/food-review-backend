package food

// Service represents for food service
type Service interface {
	AddFood(food ...Food) error
	ListAllFood() []Food
}

// NewService creates a new service
func NewService(repo Repository) Service {
	return &service{repo}
}

type service struct {
	repository Repository
}

func (s *service) AddFood(food ...Food) error {
	for _, f := range food {
		_, err := s.repository.Create(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) ListAllFood() []Food {
	return s.repository.GetAll()
}
