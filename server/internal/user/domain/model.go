package domain

type User struct {
	ID       string
	Email    string
	Password string // hashed
}

func (u *User) CheckPassword(rawPassword string) bool {
	return u.Password == HashPassword(rawPassword) // simple
}

func HashPassword(p string) string {
	return "hashed_" + p // mock hash
}
