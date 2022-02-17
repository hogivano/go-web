package user

type LoginUserInput struct {
	Email    string
	Password string
}

type RegisterUserInput struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secretKey"`
}
