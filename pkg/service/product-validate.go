package service

import "github.com/aashpv/backend-microservice/pkg/models"

func (s *service) ValidateProduct(product models.Product) (err error) {
	err = s.v.Struct(product)
	if err != nil {
		return
	}
	return
}
