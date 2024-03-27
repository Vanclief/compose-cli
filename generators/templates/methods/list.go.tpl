package {{.PackageName}}

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vanclief/ez"

	"{{.ModulePath}}/application/models/{{.PackageName}}"

	"github.com/vanclief/compose/interfaces/rest/requests"
	"github.com/vanclief/compose/interfaces/rest/responses"
)

type ListRequest struct {
	requests.StandardList
    Search      string `json:"search"`
}

func (r ListRequest) Validate() error {
	const op = "ListRequest.Validate"

	err := r.StandardList.Validate()
	if err != nil {
		return ez.Wrap(op, err)
	}

	err = validation.ValidateStruct(&r,
		// validation.Field(&r.Field, validation.Required),
	)

	if err != nil {
		return ez.New(op, ez.EINVALID, err.Error(), nil)
	}

	return nil
}

type ListResponse struct {
	responses.StandardList
	{{.ModelSlice}} []models.{{.ModelStruct}} `json:"{{.PackageName}}"`
}

func (api *API) List(ctx context.Context, requester *models.User, request *ListRequest) (*ListResponse, error) {
	const op = "API.List"

    var count int
	{{.PackageName}}List := []models.{{.ModelStruct}}{}

	response := &ListResponse{
		StandardList: responses.StandardList{
			Limit:      request.Limit,
			Offset:     request.Offset,
			TotalCount: count,
		},
		{{.ModelSlice}}: {{.PackageName}}List,
	}

	err := response.GenerateHash(response.{{.ModelSlice}})
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil
}
