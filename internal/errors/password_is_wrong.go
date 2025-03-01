package errors

type PasswordIsWrong struct{}

func (PasswordIsWrong) Error() string {
	return "Password Is Wrong"
}
