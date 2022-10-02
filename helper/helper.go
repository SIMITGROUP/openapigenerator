package helper

import (
	"github.com/getkin/kin-openapi/openapi3"
)

/*
	 	5 allresponse
		6 all requestbodies
*/

func PrepareObjects(document *openapi3.T) {
	Doc = document
	//prepare Allschemas
	PrepareSchemas()
	// prepare AllSecuritySchemes
	PrepareSecuritySchemes()
	//prepare Allroutes, AllRequestHandles,
	PrepareRoutes()
}
