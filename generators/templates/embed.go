package templates

import "embed"

//go:embed api.go.tpl
//go:embed api_test.go.tpl
var FS embed.FS
