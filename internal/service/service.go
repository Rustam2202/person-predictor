package service

type Service struct {
	PersonService *PersonService
}

func NewService(pr PersonRepository) *Service {
	return &Service{PersonService: NewPersonService(pr)}
}
