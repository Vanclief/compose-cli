package {{.PackageName}}

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vanclief/ez"

	"{{.ModulePath}}/application/resources/models/{{.ModelPackage}}"
)

type {{.MethodName}}Request struct {
	{{.ModelStruct}}ID     int64  `json:"{{.ModelVariable}}_id"`
}


func (r {{.MethodName}}Request) Validate() error {
	const op = "{{.MethodName}}Request.Validate"

	err := validation.ValidateStruct(&r,
		validation.Field(&r.{{.ModelStruct}}ID, validation.Required),
	)
	if err != nil {
		return ez.New(op, ez.EINVALID, err.Error(), nil)
	}

	return nil
}

func (api *API) {{.MethodName}}(ctx context.Context, requester *models.User, request *{{.MethodName}}Request) (int64, error) {
	const op = "API.{{.MethodName}}"

	return 0, nil
}