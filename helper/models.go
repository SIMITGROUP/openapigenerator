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
	Path        string
	Method      string
	OperationID string
	Middlewares []string
	Summary     string
	Description string
	Responses   openapi3.Responses
	DataType    string
}
type RouteSettings struct {
	Methods []MethodSettings
}

type Model_Handle struct {
	FuncName string
	DataType string
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
