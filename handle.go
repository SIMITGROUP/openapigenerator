package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func stopWhenNoOperationId(pathstr string, methodname string, operationID string) {
	if operationID == "" {
		log.Fatal(methodname + "  '" + pathstr + "' does not define operationId")
	}
}

func prepareHandles(doc *openapi3.T) {

	var route_executors []string
	executor_result := make(map[string]string)
	handlestring := ""
	// allschema := doc.Components.Schemas

	for pathstr, pathmethods := range doc.Paths {
		if pathmethods.Get != nil {
			executor := pathmethods.Get.OperationID
			// stopWhenNoOperationId(pathstr, "GET", executor)
			identifiedUnsupportedMethod(pathstr, "GET", pathmethods.Get)
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Get.Responses)

		}
		if pathmethods.Put != nil {
			executor := pathmethods.Put.OperationID
			identifiedUnsupportedMethod(pathstr, "PUT", pathmethods.Put)
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Put.Responses)
		}
		if pathmethods.Post != nil {
			executor := pathmethods.Post.OperationID
			identifiedUnsupportedMethod(pathstr, "Post", pathmethods.Post)
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Post.Responses)
		}
		if pathmethods.Delete != nil {
			executor := pathmethods.Delete.OperationID
			identifiedUnsupportedMethod(pathstr, "Delete", pathmethods.Delete)
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Delete.Responses)
		}

		route_executors = removeDuplicate(route_executors)
	}

	if len(route_executors) > 0 {

		for _, executor := range route_executors {
			// Data_funcMap = Data_funcMap + fmt.Sprintf("\n    \"%v\":%v,", executor, executor)

			handlestring = handlestring +
				fmt.Sprintf("\nfunc %v(c *gin.Context) {\n"+
					"    data := %v\n"+
					"    c.JSON(http.StatusOK, data)"+
					"\n}\n", executor, executor_result[executor])
		}
	}

	filename := "handles.go"
	content := "package openapi\n\n" +
		"import (\n" +
		"\n   \"github.com/gin-gonic/gin\"\n" +
		"\n   \"net/http\"\n)\n" +
		handlestring

	// _ = os.Remove(filename)
	// _ = os.WriteFile(filename, []byte(content), 0644)
	writeFile("openapi", filename, content)
}
func identifiedUnsupportedMethod(pathstr string, methodname string, op *openapi3.Operation) {

	if op.OperationID == "" {
		log.Fatal(methodname + "  '" + pathstr + "' does not define operationId")
	}

	for statuscode, setting := range op.Responses {
		if statuscode == "200" {
			content := setting.Value.Content["application/json"]
			_ = content
			// if content.Schema.Value.Ref {
			// 	log.Fatal(methodname, " ", pathstr, "status '200' found schema using 'oneof' which is not supported")
			// }

		}

	}

}
func getResponses(res openapi3.Responses) string {
	result := ""
	nodata := false
	examples := "{}"
	for statuscode, setting := range res {
		content := setting.Value.Content["application/json"]

		nodata = false
		// only return first 1, usually http status 200
		if statuscode == "200" {
			// fmt.Println("status:", statuscode)
			//if is reference
			if content != nil {
				if content.Schema.Ref != "" {
					if content.Schema.Value.Type == "object" {
						values := strings.Split(content.Schema.Ref, "/")
						result = GetModelName(values[len(values)-1]) // get Model name
						examples = getExamples(content.Schema.Value)

					} else if content.Schema.Value.Type == "array" {
						values := strings.Split(content.Schema.Value.Items.Ref, "/")
						model := GetModelName(values[len(values)-1])
						examples = "[]" + model + "{}" //result
						result = ""
					} else {
						// nodata = true
					}
				}
			} else {
				return `gin.H{"msg": "undefined content in api"}`
			}

		}
	}

	if nodata == true {
		return `gin.H{"msg": "undefined type"}`
	} else {
		return result + examples
	}

}
func getExamples(schema *openapi3.Schema) string {
	examplestr := ""
	prefix := "        "

	for field, setting := range schema.Properties {
		fieldname := upperCaseFirst(field)
		// fmt.Println(prefix, field, setting.Value.Type, setting.Value.Example)
		data := setting.Value.Example
		if data == nil {
			continue
		}
		tmp := ""
		switch setting.Value.Type {
		case "string":
			tmp = fmt.Sprintf("%v: \"%v\",\n", fieldname, setting.Value.Example)
		case "array":
			tmp = fmt.Sprintf("%v: %#v,\n", fieldname, setting.Value.Example)
			tmp = strings.Replace(tmp, "interface {}", setting.Value.Items.Value.Type, -1)
		default:
			tmp = fmt.Sprintf("%v: %v,\n", fieldname, setting.Value.Example)
		}

		examplestr = examplestr + prefix + tmp
	}

	return "{\n" + examplestr + prefix + "}"
}
func getFieldTypeSettings(setting *openapi3.Schema) (string, string) {
	fieldtype := setting.Type
	fieldformat := setting.Format
	return fieldtype, fieldformat
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// func helloHandler(c *gin.Context) {
// 	claims := jwt.ExtractClaims(c)
// 	user, _ := c.Get(identityKey)
// 	c.JSON(200, gin.H{
// 		"userID":   claims[identityKey],
// 		"userName": user.(*User).UserName,
// 		"text":     "Hello World.",
// 	})
// }
