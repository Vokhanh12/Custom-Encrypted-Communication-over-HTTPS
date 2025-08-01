package domain

type User struct {
	ID       string
	Email    string
	Password string
}

func (u *User) CheckPassword(rawPassword string) bool {
	return u.Password == HashPassword(rawPassword)
}

func HashPassword(p string) string {
	return "hashed_" + p
}
