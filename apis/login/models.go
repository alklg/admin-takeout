package login

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
