package dtos

type LoginRequestDTO struct {
	Email    string
	Password string
}

type LoginResponseDTO struct {
	AccessToken  string
	RefreshToken string
}
