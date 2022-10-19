package helper

import "github.com/getkin/kin-openapi/openapi3"

var Proj = Model_ProjectSetting{}

var AllSchemas = make(map[string]Model_SchemaSetting)
var AllRequestHandles = make(map[string]Model_RequestHandle)
var AllSecuritySchemes = make(map[string]Model_SecuritySchemaSetting)
var AllRoutes = make(map[string]Model_Routes)
var AllFunctionName = []string{}

var Doc *openapi3.T
