package {{.PluralLower}}

import (
	"github.com/vanclief/go-unityqr-backend/application/notificator"
	"github.com/vanclief/go-unityqr-backend/application/resources/shared"
	"github.com/vanclief/go-unityqr-backend/controller"
	"github.com/vanclief/go-unityqr-backend/interfaces/database"
)

type API struct {
	DB          *database.DB
	SharedAPI   *shared.API
}

func New(ctrl *controller.Controller, sharedAPI *shared.API) *API {
	if ctrl == nil {
		panic("Controller reference is nil")
	} else if sharedAPI == nil {
		panic("SharedAPI reference is nil")
	} else if n == nil {
		panic("Notificator reference is nil")
	}

	api := &API{
		DB:          ctrl.DB,
		Notificator: n,
		SharedAPI:   sharedAPI,
	}

	return api
}
