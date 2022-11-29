package helper

// project info
type Model_ProjectSetting struct {
	GenerateFolder string
	ProjectName    string
	ApiFile        string
	ListenPort     string
	BuildLang      string
	OverrideHandle bool
	InitFunctions  []string
	AllEnvVars     map[string]string
}

// schema infos
type Model_SchemaSetting struct {
	ModelName     string //schema1 => Model_schema1
	ModelType     string //object,array
	ModelFormat   string //format if type = string, example date, binary
	InterfaceName string //Interface_schema1
	Description   string
	FieldList     map[string]Model_Field
}

type Model_Field struct {
	ModelName     string
	FieldName     string
	FieldType     string
	ChildItemType string //when type = array, child item type important
	ApiFieldType  string
	FieldFormat   string
	ApiFieldName  string
	Required      bool
	Description   string
	Example       string
	FieldIsModel  bool
}
type Model_Header struct {
	Name        string
	Type        string
	Description string
	Example     interface{}
}

// routing infos
type Model_Routes struct {
	Path            string                          // example: /user
	RequestSettings map[string]Model_RequestSetting // "get"=setting1, "post"=setting2

}

type Model_RequestSetting struct {
	Summary       string
	Description   string
	Path          string
	Method        string
	RequestHandle Model_RequestHandle                    //define handlename for this request
	Securities    map[string]Model_SecuritySchemaSetting //define what security middlewares
	// RequestBodies datatype	// route registry seems not required
	// Parameters 	datatype	// route registry seems not required
}

type Model_Responses struct {
	ContentType    string
	ResponseSchema Model_SchemaSetting
}

// routing handles info
type Model_RequestHandle struct {
	HandleName     string
	Summary        string
	Description    string
	ResponseSchema Model_SchemaSetting
	ResponseType   string
	RequestBodies  Model_RequestBody
	Parameters     map[string]Model_Parameter
	Headers        []Model_Header
	HttpStatusCode int
	ContentType    string
}

// route's requestbody (collection of info submit by client)
type Model_RequestBody struct {
	Name          string
	Description   string
	Required      bool
	RequestSchema Model_SchemaSetting
}

// route's parameters
type Model_Parameter struct {
	StoreIn         string
	Description     string
	Required        bool
	Deprecated      bool
	AllowEmptyValue bool
	Example         string
}

// openapi3.SecurityScheme
type Model_SecuritySchemaSetting = struct {
	Name        string
	SchemeName  string
	Type        string
	Description string
	In          string
	Scheme      string
	Scopes      map[string]string
}

// security schemes info, seems use original will
// type Model_SecuritySchemaSetting struct {
// 	openapi3.SecurityScheme
// }
