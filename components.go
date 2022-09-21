package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var componentlist = map[string]string{
	"callbacks":       "",
	"examples":        "",
	"headers":         "",
	"links":           "",
	"parameters":      "",
	"requestbodies":   "",
	"responses":       "",
	"schemas":         "",
	"securityschemes": "",
}

func prepareComponent(doc *openapi3.T) {

	//callback
	if doc.Components.Callbacks != nil {
		componentlist["callback"] = prepareCallbacks(doc.Components.Callbacks)
	}
	//examples
	if doc.Components.Examples != nil {
		componentlist["examples"] = prepareExamples(doc.Components.Examples)
	}
	//headers
	if doc.Components.Headers != nil {
		componentlist["headers"] = prepareHeaders(doc.Components.Headers)
	}
	//links
	if doc.Components.Links != nil {
		componentlist["links"] = prepareLinks(doc.Components.Links)
	}
	//parameters
	if doc.Components.Parameters != nil {
		componentlist["parameters"] = prepareParameters(doc.Components.Parameters)
	}

	//RequestBodies
	if doc.Components.RequestBodies != nil {
		componentlist["requestbodies"] = prepareRequestBodies(doc.Components.RequestBodies)
	}

	//Responses
	if doc.Components.Responses != nil {
		componentlist["responses"] = prepareResponses(doc.Components.Responses)
	}

	//schemas
	if doc.Components.Schemas != nil {
		componentlist["schemas"] = prepareSchemas(doc.Components.Schemas)
	}

	//securityschemes
	if doc.Components.SecuritySchemes != nil {
		componentlist["securityschemes"] = prepareSecuritySchemes(doc.Components.SecuritySchemes)
	}
	// for k, v := range componentlist {
	// 	filename := k + ".go"
	// 	_ = os.Remove(filename)
	// 	fmt.Println(filename)

	// 	if v != "" {
	// 		fmt.Println(v)
	// 		content := "package openapi\n" + v
	// 		// writeFile(filename, content)
	// 		// _ = os.WriteFile(filename, []byte(content), 0644)
	// 		writeFile("openapi", filename, content)
	// 	}
	// }

}
