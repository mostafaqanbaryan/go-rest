package entities

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	password string
}

func (u *User) Validate() error {
	return nil
}

func (u User) IsPasswordCorrect(password string) bool {
	return u.password == password
}
