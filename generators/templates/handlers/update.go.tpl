func (h *Handler) Update{{.ModelStruct}}(c echo.Context) error {
	const op = "Handler.Update{{.ModelStruct}}"

	request := requests.New(c.Request().Header, c.RealIP())

	resourceID, err := h.GetParameterID(c, "id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.UpdateRequest{
		{{.ModelStruct}}ID: resourceID,
	}

	return h.BindedRequest(c, op, request, requestBody)
}

