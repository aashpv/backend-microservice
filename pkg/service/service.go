package service

type Service interface {
	CreateProduct(body string) (err error)
	UpdateProductById(body string) (err error)
	GetProductById(id int) (err error)
	GetAllProducts() (err error)
	DeleteProductById(id int) (err error)
	CreateOrder(body string) (err error)
	UpdateOrderById(body string) (err error)
	GetOrderById(id int) (err error)
	GetAllOrders() (err error)
	DeleteOrderById(id int) (err error)
	CreateUser(body string) (err error)
	UpdateUserById(body string) (err error)
	GetUserById(id int) (err error)
	GetAllUsers() (err error)
	DeleteUserById(id int) (err error)
}

type service struct {
}
