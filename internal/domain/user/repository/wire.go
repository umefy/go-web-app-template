package repository

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(Repository), new(*userRepository)),
)
