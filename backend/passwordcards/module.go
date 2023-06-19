package passwordcards

import (
	"github.com/google/wire"
)

var PasswordCardsModule = wire.NewSet(
	ProvideController,
	ProvideService,
	ProvideRepository,
)
