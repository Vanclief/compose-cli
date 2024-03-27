package templates

import "embed"

//go:embed api.go.tpl
//go:embed api_test.go.tpl
//go:embed create.go.tpl
//go:embed delete.go.tpl
//go:embed generic.go.tpl
//go:embed generic_test.go.tpl
//go:embed get.go.tpl
//go:embed list.go.tpl
//go:embed update.go.tpl
var FS embed.FS
