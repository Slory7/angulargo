package contracts

type ResultStatus int

const (
	Success ResultStatus = (iota + 1)
	NotFound
	Forbidden
	Unauthorized
	Conflict
	BadData
	BadLogic
	Error
)

func (status ResultStatus) String() string {
	names := []string{
		"Success",
		"NotFound",
		"Forbidden",
		"Unauthorized",
		"Conflict",
		"BadData",
		"BadLogic",
		"Error",
	}
	return names[int(status)-1]
}
