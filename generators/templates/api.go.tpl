package {{.PackageName}}

import (
	"{{.ModulePath}}/application/resources/shared"
	"{{.ModulePath}}/controller"
	"{{.ModulePath}}/interfaces/database"
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
	} 

	api := &API{
		DB:          ctrl.DB,
		SharedAPI:   sharedAPI,
	}

	return api
}
