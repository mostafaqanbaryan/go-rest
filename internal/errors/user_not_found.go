package errors

type UserNotFound struct{}

func (UserNotFound) Error() string {
	return "User Not Found"
}
