package urls

import (
	"github.com/slory7/angulargo/src/infrastructure/business/constants"
)

type QueryObject struct {
	Start       int
	Limit       int
	OrderBy     string
	IsDecending bool
	FilterBy    string
	Op          constants.Operator
	FilterValue string
}
