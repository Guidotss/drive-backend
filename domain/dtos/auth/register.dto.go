package auth

type RegisterDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	Name     string `json:"name" validate:"required,min=3,max=50"`
}
