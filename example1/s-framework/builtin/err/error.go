package err

type Error struct {
	Code  string
	Cause string
}

func (e Error) Error() string {
	return ""
}
