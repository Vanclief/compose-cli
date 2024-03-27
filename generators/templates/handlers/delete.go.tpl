func (h *Handler) Delete{{.ModelStruct}}(c echo.Context) error {
	const op = "Handler.Delete{{.ModelStruct}}"

	request := requests.New(c.Request().Header, c.RealIP())

	resourceID, err := h.GetParameterID(c, "id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.DeleteRequest{
		{{.ModelStruct}}ID: resourceID,
	}

	return h.StandardRequest(c, op, request, requestBody)
}

