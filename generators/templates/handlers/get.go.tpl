func (h *Handler) Get{{.ModelStruct}}(c echo.Context) error {
	const op = "Handler.Get{{.ModelStruct}}"

	request := requests.New(c.Request().Header, c.RealIP())

	resourceID, err := h.GetParameterID(c, "id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.GetRequest{
		{{.ModelStruct}}ID: resourceID,
	}

	return h.StandardRequest(c, op, request, requestBody)
}

