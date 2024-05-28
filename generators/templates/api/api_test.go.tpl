package {{.PackageName}}

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"{{.ModulePath}}/application/resources/shared/test"
)

type {{.SuiteName}} struct {
	suite.Suite
	test.TestSuite
	api *API
}

func (suite *{{.SuiteName}}) SetupTest() {
	suite.Setup()

	suite.api = New(suite.Ctrl, suite.SharedAPI)
}

func (suite *{{.SuiteName}}) TearDownTest() {
	suite.TCtrl.Teardown()
}

func Test{{.SuiteName}}(t *testing.T) {
	suite.Run(t, new({{.SuiteName}}))
}
