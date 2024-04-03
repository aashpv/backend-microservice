package service

import "github.com/aashpv/backend-microservice/pkg/models"

func (s *service) ValidateUser(user models.User) (err error) {
	err = s.v.Struct(user)
	if err != nil {
		return
	}
	return
}
