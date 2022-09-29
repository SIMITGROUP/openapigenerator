package helper

import "github.com/getkin/kin-openapi/openapi3"

type ProjectSetting struct {
	GenerateFolder string
	ProjectName    string
	ApiFile        string
	ListenAddress  string
	BuildLang      string
}

type ProjectPara struct {
	Project string
	Listen  string
}
type MethodSettings struct {
	Path              string
	Method            string
	OperationID       string
	Middlewares       []string
	Summary           string
	Description       string
	Responses         openapi3.Responses
	DataType          string
	RequestBodiesName string
}
type RouteSettings struct {
	Methods []MethodSettings
}

type Model_Handle struct {
	FuncName   string
	DataType   string
	SchemaType string
	// Example  string
}
type Model_HandleTemplate struct {
	Handles []Model_Handle
}

type Model_Field struct {
	ModelName    string
	FieldName    string
	FieldType    string
	ApiFieldName string
	Description  string
	Example      string
}

type Model_Schema struct {
	ModelName     string
	InterfaceName string
	FieldList     []Model_Field
}

type Model_Security struct {
	Name           string
	ModelName      string
	InterfaceName  string
	MethodName     string
	HandleName     string
	HandleData     string
	SecurityType   string
	SecurityScheme string
}
