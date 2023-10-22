package persons

import "person-predicator/internal/service"

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(s *service.PersonService) *PersonHandler {
	return &PersonHandler{service: s}
}
