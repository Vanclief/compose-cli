package templates

import "embed"

//go:embed api/api.go.tpl
//go:embed api/api_test.go.tpl

//go:embed methods/create.go.tpl
//go:embed methods/delete.go.tpl
//go:embed methods/generic.go.tpl
//go:embed methods/generic_test.go.tpl
//go:embed methods/get.go.tpl
//go:embed methods/list.go.tpl
//go:embed methods/update.go.tpl

//go:embed handlers/create.go.tpl
//go:embed handlers/delete.go.tpl
//go:embed handlers/generic.go.tpl
//go:embed handlers/get.go.tpl
//go:embed handlers/imports.go.tpl
//go:embed handlers/list.go.tpl
//go:embed handlers/update.go.tpl

var FS embed.FS
