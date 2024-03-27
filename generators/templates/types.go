package templates

type APIData struct {
	PackageName string
	ModulePath  string
}

type APITestData struct {
	PackageName string
	ModulePath  string
	SuiteName   string
}

type MethodData struct {
	PackageName   string // codes
	ModulePath    string // go.mod path
	ModelStruct   string // Code
	ModelVariable string // code
	ModelSlice    string // Codes
	MethodName    string // Create
}

type MethodTestData struct {
	PackageName string
	ModulePath  string
	SuiteName   string
	TestFunc    string
}
