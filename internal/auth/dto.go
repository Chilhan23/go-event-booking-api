package auth

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}


type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct{
	Token string `json:"token"`
	User UserResponse `json:"user"`
}

type UserResponse struct{
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}