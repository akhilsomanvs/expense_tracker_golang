package entities

type SignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"raquired,email"`
	Password string `json:"password" validate:"required,min=8"`
}
