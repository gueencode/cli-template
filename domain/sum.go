package domain

import (
	"github.com/gueencode/gwcli/domain/model"
	"github.com/gueencode/gwcli/util"
)

// NewNumbersFromStringSlice create new Numbers with numbers from string slice
func NewNumbersFromStringSlice(strNumbers []string) (model.Numbers, error) {
	rawNumbers, err := util.ConvertStringSliceToIntSlice(strNumbers)
	if err != nil {
		return nil, err
	}
	return model.NewNumbers(rawNumbers), nil
}
