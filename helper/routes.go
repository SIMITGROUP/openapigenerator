package helper

import (
	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

// consolidate all routes and route's handles
func PrepareRoutes() {

	for path, pathmethods := range Doc.Paths {

		req := make(map[string]Model_RequestSetting)
		if pathmethods.Get != nil {
			req["GET"] = getPathSetting("GET", path, pathmethods.Get)
		}
		if pathmethods.Post != nil {
			req["POST"] = getPathSetting("POST", path, pathmethods.Post)
		}
		if pathmethods.Put != nil {
			req["PUT"] = getPathSetting("PUT", path, pathmethods.Put)
		}
		if pathmethods.Delete != nil {
			req["DELETE"] = getPathSetting("DELETE", path, pathmethods.Delete)
		}
		if pathmethods.Head != nil {
			req["HEAD"] = getPathSetting("HEAD", path, pathmethods.Head)
		}
		if pathmethods.Patch != nil {
			req["PATCH"] = getPathSetting("PATCH", path, pathmethods.Patch)
		}
		if pathmethods.Options != nil {
			req["OPTIONS"] = getPathSetting("OPTIONS", path, pathmethods.Options)
		}
		if pathmethods.Trace != nil {
			req["TRACE"] = getPathSetting("TRACE", path, pathmethods.Trace)
		}
		AllRoutes[path] = Model_Routes{
			Path:            path,
			RequestSettings: req,
		}
	}
}

func getPathSetting(methodtype string, path string, op *openapi3.Operation) Model_RequestSetting {
	log.Info(methodtype, " ", path)
	rsetting := Model_RequestSetting{
		Summary:       op.Summary,
		Path:          path,
		Method:        methodtype,
		Description:   op.Description,
		Securities:    GetSecurityMiddleware(methodtype, path, op),
		RequestHandle: GetRequestHandle(methodtype, path, op),
	}
	return rsetting
}

func GetRequestHandle(methodtype string, path string, op *openapi3.Operation) Model_RequestHandle {

	if op.OperationID == "" {
		log.Fatal(methodtype, path, " undefine operationId schema RequestBodies")
	}
	log.Info("    Handle: ", op.OperationID)

	requestobj, ok := AllRequestHandles[op.OperationID]
	if ok {
		//do nothing
		return requestobj
	} else {

		//get this handle return data (object)
		responseschema := GetResponseSchema(methodtype, path, op)
		//get this handle request body (object)
		requestbody := GetRequestBodySetting(methodtype, path, op)
		//get this handle parameters (array)
		paras := GetParameters(methodtype, path, op)
		handle := Model_RequestHandle{
			HandleName:     op.OperationID,
			ResponseSchema: responseschema,
			Parameters:     paras,
			RequestBodies:  requestbody,
		}
		AllRequestHandles[op.OperationID] = handle
		return handle
	}
}
func GetResponseSchema(methodtype string, path string, op *openapi3.Operation) Model_SchemaSetting {
	selectedstatus := 200
	schema := Model_SchemaSetting{}
	selectedmime := "application/json"
	contentSchema := op.Responses.Get(selectedstatus).Value.Content.Get(selectedmime).Schema
	if contentSchema.Ref != "" {
		log.Info("        response schema name: ", contentSchema.Ref)
		schemaname := GetTypeNameFromRef(contentSchema.Ref)
		schema = AllSchemas[schemaname]
	} else {
		log.Fatal("        undefined response schema $ref")
	}

	return schema
}
func GetRequestBodySetting(methodtype string, path string, op *openapi3.Operation) Model_RequestBody {
	log.Info("        request body: ", op.RequestBody)
	requestBody := Model_RequestBody{}

	//op.RequestBody.Ref
	if op.RequestBody != nil {

		content := op.RequestBody.Value.Content
		if content != nil && content.Get("application/json").Schema.Ref != "" {
			ref := content.Get("application/json").Schema.Ref
			log.Info("        request body schema name: ", ref)

			requestBody.Description = op.RequestBody.Value.Description
			requestBody.Required = op.RequestBody.Value.Required
			schemaname := GetTypeNameFromRef(ref)

			requestBody.RequestSchema = AllSchemas[schemaname]

		} else {
			log.Fatal("        request body undefine $ref")
		}

	} // else {
	// 	log.Fatal("        undefined response schema $ref")
	// }

	return requestBody

}

func GetSecurityMiddleware(methodtype string, path string, op *openapi3.Operation) map[string]Model_SecuritySchemaSetting {

	securityrules := map[string]Model_SecuritySchemaSetting{}
	if op.Security != nil {
		log.Info("    Security middleware: ")
		for _, securitysetting := range *op.Security {
			for authname, authsetting := range securitysetting {
				/*limitation authsetting in path temporary ignore*/
				securityrules[authname] = AllSecuritySchemes[authname]

				log.Info("        ", authname)
				_ = authsetting
			}
		}
	}
	return securityrules
}
func GetParameters(methodtype string, path string, op *openapi3.Operation) map[string]Model_Parameter {
	log.Info("        parameters: ")
	paras := map[string]Model_Parameter{}
	for _, psetting := range op.Parameters {
		pname := psetting.Value.Name
		ptype := psetting.Value.Schema.Value.Type
		prequired := psetting.Value.Required
		pstorein := LowerCaseFirst(psetting.Value.In)
		if pstorein == "cookie" {
			log.Fatal("parameter ", pname, " store in cookie which is not supported")
		}

		if VerifyKeyname(pname) == false {
			log.Fatal("Invalid parameter ", pname, ", it should only consist character a-z without special character and spacing")
		}
		log.Info("            ", pname, ": ", ptype,
			", IN: ", pstorein,
			", Required: ", prequired)

		paras[psetting.Value.Name] = Model_Parameter{
			StoreIn:         psetting.Value.In,
			Required:        psetting.Value.Required,
			Deprecated:      psetting.Value.Deprecated,
			AllowEmptyValue: psetting.Value.AllowEmptyValue,
		}
		// } else {
		// psetting.Value
		// }

	}
	return paras
}

// prepare handles of each request type
// also consolidate all route's request handle (function) into single registry (Allhandles)

// func getRouteSetting(methodtype string, path string, op *openapi3.Operation) Model_RouteSetting {
// 	securityrules := map[string]Model_SecuritySchemaSetting{}
// 	if op.OperationID == "" {
// 		log.Fatal(methodtype, " ", path, " undefine operationId")
// 	}
// 	if op.Security == nil {
// 		//no midleware
// 	} else {

// 		for _, securitysetting := range *op.Security {
// 			for authname, _ := range securitysetting {

// 				/*limitation not yet support securities setting in path setting yet*/
// 				securityrules[authname] = Model_SecuritySchemaSetting{}
// 			}
// 		}

// 	}

// 	schemaobj := GetSchemaFromRef(op.RequestBody.Ref)

// 	rsetting := Model_RouteSetting{
// 		Summary:     op.Summary,
// 		Description: op.Description,
// 		RouteHandle: Model_RouteHandleSetting{
// 			HandleName:     op.OperationID,
// 			ResponseSchema: schemaobj,
// 		},
// 		Securities: securityrules,
// 	}
// 	paras := map[string]Model_Parameter{}
// 	for _, psetting := range op.Parameters {
// 		if psetting.Ref == "" {
// 			paras[psetting.Value.Name] = Model_Parameter{
// 				StoreIn:         psetting.Value.In,
// 				Required:        psetting.Value.Required,
// 				Deprecated:      psetting.Value.Deprecated,
// 				AllowEmptyValue: psetting.Value.AllowEmptyValue,
// 			}
// 		}

// 	}

// 	rsetting.RouteHandle.Parameters = paras
// 	return rsetting
// }
