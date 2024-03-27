package {{.PackageName}}

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vanclief/ez"

	"{{.ModulePath}}/application/models"
)

type {{.MethodName}}Request struct {}

func (r {{.MethodName}}Request) Validate() error {
	const op = "{{.MethodName}}Request.Validate"

	err := validation.ValidateStruct(&r,
		// validation.Field(&r.Field, validation.Required),
	)
	if err != nil {
		return ez.New(op, ez.EINVALID, err.Error(), nil)
	}

	return nil
}

func (api *API) {{.MethodName}}(ctx context.Context, requester *models.User, request *{{.MethodName}}Request) (*models.{{.ModelStruct}}, error) {
	const op = "API.{{.MethodName}}"

	return nil, nil
}
