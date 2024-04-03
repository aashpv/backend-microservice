package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Number   string `json:"number" validate:"e164, required"`
	Email    string `json:"email" validate:"email, required"`
}

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
	CategoryID  int     `json:"id_category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}

type Order struct {
	ID          int    `json:"id"`
	Date        string `json:"date" validate:"datetime, required"`
	Status      string `json:"status" validate:"required"`
	ProductID   int    `json:"id_product" validate:"required"`
	UserID      int    `json:"id_user" validate:"required"`
	Address     string `json:"address" validate:"required"`
	OtherNumber string `json:"other_number" validate:"e164"`
	OtherName   string `json:"other_name"`
}
