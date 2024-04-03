package service

import (
	"github.com/aashpv/backend-microservice/pkg/models"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	//CreateProduct(product models.Product) (err error)
	//UpdateProductById(product models.Product) (err error)
	//GetProductById(id int) (product models.Product, err error)
	//GetAllProducts() (product models.Product, err error)
	//DeleteProductById(id int) (err error)
	//CreateOrder(order models.Order) (err error)
	//UpdateOrderById(order models.Order) (err error)
	//GetOrderById(id int) (order models.Order, err error)
	//GetAllOrders() (order models.Order, err error)
	//DeleteOrderById(id int) (err error)
	//CreateUser(user models.User) (err error)
	//UpdateUserById(user models.User) (err error)
	//GetUserById(id int) (user models.User, err error)
	//GetAllUsers() (user models.User, err error)
	//DeleteUserById(id int) (err error)
	ValidateOrder(order models.Order) (err error)
	ValidateProduct(product models.Product) (err error)
	ValidateUser(user models.User) (err error)
}

type service struct {
	v validator.Validate
}

func New(validator validator.Validate) Service {
	return &service{v: validator}
}
