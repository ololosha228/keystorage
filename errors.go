package keystorage

type ErrUserNotFound struct {
	User string
}

func (e *ErrUserNotFound) Error() string {
	return "username '" + e.User + "' not found"
}
