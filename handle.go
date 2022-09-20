package main

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func prepareHandles(doc *openapi3.T) {
	var route_executors []string
	executor_result := make(map[string]string)
	handlestring := ""
	// allschema := doc.Components.Schemas

	for _, pathmethods := range doc.Paths {
		if pathmethods.Get != nil {
			executor := pathmethods.Get.OperationID
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Get.Responses)

		}
		if pathmethods.Put != nil {
			executor := pathmethods.Put.OperationID
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Put.Responses)
		}
		if pathmethods.Post != nil {
			executor := pathmethods.Post.OperationID
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Post.Responses)
		}
		if pathmethods.Delete != nil {
			executor := pathmethods.Delete.OperationID
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

func getResponses(res openapi3.Responses) string {
	result := ""
	examples := "{}"
	for statuscode, setting := range res {
		content := setting.Value.Content["application/json"]

		// only return first 1, usually http status 200
		if content != nil && statuscode == "200" {
			fmt.Println("status:", statuscode)
			values := strings.Split(content.Schema.Ref, "/")
			result = values[len(values)-1]
			examples = getExamples(content.Schema.Value)
			break
		}
	}

	if result == "" {
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
