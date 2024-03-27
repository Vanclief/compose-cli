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
	PackageName   string
	ModulePath    string
	ModelPackage  string
	ModelStruct   string
	ModelVariable string
	ModelSlice    string
	MethodName    string
}

type MethodTestData struct {
	PackageName string
	ModulePath  string
	SuiteName   string
	TestFunc    string
}
