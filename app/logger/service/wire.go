package service

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewService,
	wire.Bind(new(Service), new(*loggerService)),
)
