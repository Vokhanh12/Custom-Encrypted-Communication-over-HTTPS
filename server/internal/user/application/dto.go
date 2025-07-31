package application

type LoginRequestDTO struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	Token string
}
