//go:build wireinject
// +build wireinject

package registry

//go:generate wire

import (
	"github.com/google/wire"
	"github.com/gueencode/gwcli/domain/model"
	"github.com/gueencode/gwcli/handler"
	"github.com/gueencode/gwcli/infra"
	"github.com/gueencode/gwcli/infra/repoimpl"
	"github.com/gueencode/gwcli/usecase"
	"github.com/labstack/echo"
)

// InitializeHandler initialize handlers with memorySumHistoryRepository
func InitializeHandler(v []*model.SumHistory) *handler.Handlers {
	wire.Build(handler.New, usecase.NewSum, repoimpl.NewMemorySumHistory)
	return &handler.Handlers{}
}

// InitializeSumUseCase initialize sum use case with  memorySumHistoryRepository
func InitializeSumUseCase(v []*model.SumHistory) *usecase.Sum {
	wire.Build(repoimpl.NewMemorySumHistory, usecase.NewSum)
	return &usecase.Sum{}
}

// InitializeServer initialize echo server with memorySumHistoryRepository
func InitializeServer(v []*model.SumHistory) *echo.Echo {
	wire.Build(
		handler.New,
		repoimpl.NewMemorySumHistory,
		usecase.NewSum,
		infra.NewServer,
	)
	return &echo.Echo{}
}
