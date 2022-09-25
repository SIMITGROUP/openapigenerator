package helper

import (
	"log"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func ReadRoutes(doc *openapi3.T) {

	for oripath, pathmethods := range doc.Paths {

		path := ConvertGinPath(oripath)

		if pathmethods.Get != nil {

			methodsetting := generateMethodObject("GET", path, pathmethods.Get)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Post != nil {
			methodsetting := generateMethodObject("POST", path, pathmethods.Post)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Put != nil {
			methodsetting := generateMethodObject("PUT", path, pathmethods.Put)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Delete != nil {
			methodsetting := generateMethodObject("DELETE", path, pathmethods.Delete)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Head != nil {
			methodsetting := generateMethodObject("HEAD", path, pathmethods.Head)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Patch != nil {
			methodsetting := generateMethodObject("PATCH", path, pathmethods.Patch)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Options != nil {
			methodsetting := generateMethodObject("OPTIONS", path, pathmethods.Options)
			Allroutes = append(Allroutes, methodsetting)
		}
		if pathmethods.Trace != nil {
			methodsetting := generateMethodObject("TRACE", path, pathmethods.Trace)
			Allroutes = append(Allroutes, methodsetting)
		}
	}

}

// prepare properties of each route
func generateMethodObject(methodtype string, path string, op *openapi3.Operation) MethodSettings {
	middlewares := []string{}
	securities := op.Security
	if securities == nil {
		//no midleware
	} else {
		for _, securitysetting := range *securities {
			for authname, authsetting := range securitysetting {
				methodname := GetAuthMethodName(authname)
				_ = authsetting
				handle := methodname
				middlewares = append(middlewares, handle)
			}
		}

	}
	m := MethodSettings{
		Path:        path,
		Method:      methodtype,
		OperationID: op.OperationID,
		Middlewares: middlewares,
		Summary:     op.Summary,
		Description: strings.Replace(op.Description, "\n", "\n    //", -1),
		DataType:    GetResponseDataType(path, methodtype, op.Responses),
		Responses:   op.Responses,
	}
	return m
}

func GetRoutes() RouteSettings {
	routesettings := RouteSettings{Allroutes}
	return routesettings
}

// prepare sample response for route handle
// only capture http status 200, and content for application/json
// generate sample data according api document too
func GetResponseDataType(path string, method string, responses openapi3.Responses) string {
	datatype := ""
	for statuscode, res := range responses {
		if statuscode == "200" {
			if res.Value.Content == nil {
				log.Fatal(method, " ", path, " status '200' undefine content")
			}
			if res.Value.Content["application/json"] == nil {
				log.Fatal(method, " ", path, " status '200' undefine application/json")
			}
			content := res.Value.Content["application/json"]
			if content.Schema.Ref == "" {
				log.Fatal(method, " ", path, " status '200' undefine application/json schema.$ref")
			}
			datatype := GetModelNameFromRef(content.Schema.Ref)

			return datatype
			// properties := content.Schema.Value.Properties
			// for field, fsetting := range properties {
			// 	if fsetting.Value.Type == "object" {
			// 		values := strings.Split(content.Schema.Ref, "/")
			// 		datatype = helper.GetModelName(values[len(values)-1]) // get Model name
			// 		examples = getExamples(content.Schema.Value)
			// 	} else if fsetting.Value.Type == "array" {
			// 		values := strings.Split(content.Schema.Value.Items.Ref, "/")
			// 		datatype := helper.GetModelName(values[len(values)-1])
			// 		examples = "[]" + model + "{}" //result
			// 	} else {
			// 	}

			// }
			// fmt.Println("properties", properties)

		}

	}
	return datatype
}

func GetSchemaFromName(schemaname string) *openapi3.Schema {

	for name, setting := range Allschemas {
		if schemaname == name {

			return setting.Value
		}
	}
	dummy := openapi3.Schema{Type: ""}
	return &dummy
}
