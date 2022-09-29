package helper

import "github.com/getkin/kin-openapi/openapi3"

func ReadComponents(doc *openapi3.T) {
	// if doc.Components.Callbacks != nil {
	// 	componentlist["callback"] = prepareCallbacks(doc.Components.Callbacks)
	// }
	// //examples
	// if doc.Components.Examples != nil {
	// 	componentlist["examples"] = prepareExamples(doc.Components.Examples)
	// }
	// //headers
	// if doc.Components.Headers != nil {
	// 	componentlist["headers"] = prepareHeaders(doc.Components.Headers)
	// }
	// //links
	// if doc.Components.Links != nil {
	// 	componentlist["links"] = prepareLinks(doc.Components.Links)
	// }
	// //parameters
	// if doc.Components.Parameters != nil {
	// 	componentlist["parameters"] = prepareParameters(doc.Components.Parameters)
	// }

	//RequestBodies
	if doc.Components.RequestBodies != nil {
		Allrequestbodies = doc.Components.RequestBodies
	}

	// //Responses
	// if doc.Components.Responses != nil {
	// 	componentlist["responses"] = prepareResponses(doc.Components.Responses)
	// }

	//schemas
	if doc.Components.Schemas != nil {
		Allschemas = doc.Components.Schemas
	}

	// //securityschemes
	if doc.Components.SecuritySchemes != nil {
		Allsecurityschemas = doc.Components.SecuritySchemes
	}
}
