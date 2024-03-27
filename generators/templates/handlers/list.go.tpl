func (h *Handler) List{{.ModelSlice}}(c echo.Context) error {
	const op = "Handler.List{{.ModelSlice}}"

	request := requests.New(c.Request().Header, c.RealIP())

	parentID, err := h.GetParameterID(c, "parent_id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.ListRequest{
		ParentID: parentID,
		StandardList: requests.StandardList{
			Limit:  h.GetListLimit(c, 50),
			Offset: h.GetListOffest(c, 0),
		},
		Search: c.QueryParam("search"),
	}

	return h.StandardRequest(c, op, request, requestBody)
}

