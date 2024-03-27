func (h *Handler) Create{{.ModelStruct}}(c echo.Context) error {
	const op = "Handler.Create{{.ModelStruct}}"

	request := requests.New(c.Request().Header, c.RealIP())

	parentID, err := h.GetParameterID(c, "parent_id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.CreateRequest{
		ParentID: parentID,
	}

	return h.BindedRequest(c, op, request, requestBody)
}

