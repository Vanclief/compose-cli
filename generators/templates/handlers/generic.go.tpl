func (h *Handler) {{.MethodName}}{{.ModelStruct}}(c echo.Context) error {
	const op = "Handler.{{.MethodName}}{{.ModelStruct}}"

	request := requests.New(c.Request().Header, c.RealIP())

	resourceID, err := h.GetParameterID(c, "id")
	if err != nil {
		return h.ManageError(c, op, request, err)
	}

	requestBody := &{{.PackageName}}.{{.MethodName}}{{.ModelStruct}}{
		{{.ModelStruct}}ID: resourceID,
	}

    // Use BindedRequest if the body should be binded to the body
    // which is the case if you are passing a JSON body in the request
    // otherwise use return h.StandardRequest(c, op, request, requestBody)
	return h.BindedRequest(c, op, request, requestBody)
}

