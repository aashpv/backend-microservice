package service

import "github.com/aashpv/backend-microservice/pkg/models"

func (s *service) ValidateOrder(order models.Order) (err error) {
	err = s.v.Struct(order)
	if err != nil {
		return
	}
	return
}
