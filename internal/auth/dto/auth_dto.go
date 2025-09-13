package dto

type SignUpRequest struct {
	FullName string `json:"full_name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type ConfirmEmailRequest struct {
	UserId uint   `uri:"userId" binding:"required"`
	Token  string `uri:"token" binding:"required"`
}
