package errors

type ValueError string

func (e ValueError) Error() string {
	return "value error: " + string(e)
}
