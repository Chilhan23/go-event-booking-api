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

type UserResponse struct {
	ID       string `json:"id" example:"5b676171-999a-4d62-8506-c0841707b118"`
	Username string `json:"username" example:"johndoe"`
	Email    string `json:"email" example:"johndoe@example.com"`
}

type AuthResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserResponse `json:"user"`
}

type RegisterSuccessResponse struct {
	Message string       `json:"message" example:"user registered successfully"`
	Data    UserResponse `json:"data"`
}

type LoginSuccessResponse struct {
	Message string       `json:"message" example:"login successful"`
	Data    AuthResponse `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"invalid username or password"`
}