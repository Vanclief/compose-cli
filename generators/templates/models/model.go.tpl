package models

import (
	"github.com/uptrace/bun"
)

type {{.ModelStruct}} struct {
	bun.BaseModel `bun:"table:{{.PackageName}}"`
}

