package repository

import "github.com/mpppk/gwcli/domain/model"

// SumHistory represents repository to manage history of sum calculation
type SumHistory interface {
	Add(sumHistory *model.SumHistory)
	List(limit int) []*model.SumHistory
}
