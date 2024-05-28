package {{.PackageName}}

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vanclief/ez"

	"{{.ModulePath}}/application/models"

	"github.com/vanclief/compose/interfaces/rest/requests"
	"github.com/vanclief/compose/interfaces/rest/responses"
)

type ListRequest struct {
	requests.KeysetBasedList
    Search      string `json:"search"`
}

func (r ListRequest) Validate() error {
	const op = "ListRequest.Validate"

	err := r.KeysetBasedList.Validate()
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
	responses.KeysetBasedList
	{{.ModelSlice}} []models.{{.ModelStruct}} `json:"{{.PackageName}}"`
}

func (api *API) List(ctx context.Context, requester *models.User, request *ListRequest) (*ListResponse, error) {
	const op = "API.List"

	{{.PackageName}}List := []models.{{.ModelStruct}}{}

	response := &ListResponse{
		KeysetBasedList: responses.KeysetBasedList{
			Limit: request.Limit,
		},
		{{.ModelSlice}}: {{.PackageName}}List,
	}

	responseLength, err := response.FinalizeResponse({{.PackageName}}List, len({{.PackageName}}List))
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	{{.PackageName}}List = {{.PackageName}}List[:responseLength]
	if response.HasNextPage {
		response.NextCursor = {{.PackageName}}List[len({{.PackageName}}List)-1].GetCursor()
	}

	response.{{.ModelSlice}} = {{.PackageName}}List

	return response, nil
}
