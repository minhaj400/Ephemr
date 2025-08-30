package dto

type SignUpRequest struct {
	FullName string `json:"full_name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
