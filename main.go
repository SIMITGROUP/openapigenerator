package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("undefine spec file")
	} else {
		docfile := args[1]

		doc, _ := openapi3.NewLoader().LoadFromFile(docfile)
		readAPI(doc)
		schemafile := "openapi/schema.go"
		openapifile := "openapi/openapi.go"
		userfunctionfile := "openapi/userfunction.go"
		_ = os.WriteFile(schemafile, prepareSchema(), 0644)
		_ = os.WriteFile(openapifile, prepareApi(), 0644)
		_ = os.WriteFile(userfunctionfile, prepareUserFunction(), 0644)

		fmt.Println("microservices code generated ", schemafile, openapifile, userfunctionfile)

	}
}

func prepareSchema() []byte {
	data := strings.Replace(Temp_schema, "##data##", Data_schema, -1)
	return []byte(data)
}

func prepareApi() []byte {
	data := strings.Replace(Temp_api, "##data##", Data_api, -1)
	return []byte(data)
}

func prepareUserFunction() []byte {
	data := strings.Replace(Temp_userfunction, "##data##", Data_userfunction, -1)
	return []byte(data)
}

func getResponses(res openapi3.Responses) string {
	result := ""
	for statuscode, setting := range res {
		content := setting.Value.Content["application/json"]

		fmt.Println("get response ", statuscode)
		if content != nil {
			values := content.Schema.Value.Properties
			for fieldname, _ := range values {
				result = result + "\n        \"" + fieldname + "\": \"" + fieldname + "\","
			}
		}
	}
	return result

}

func readAPI(doc *openapi3.T) {
	var route_executors []string
	executor_result := make(map[string]string)
	// allschema := doc.Components.Schemas

	for _, pathmethods := range doc.Paths {
		if pathmethods.Get != nil {
			executor := pathmethods.Get.OperationID
			route_executors = append(route_executors, executor)
			executor_result[executor] = getResponses(pathmethods.Get.Responses)

		}
		if pathmethods.Put != nil {
			route_executors = append(route_executors, pathmethods.Put.OperationID)
		}
		if pathmethods.Post != nil {
			route_executors = append(route_executors, pathmethods.Post.OperationID)
		}
		if pathmethods.Delete != nil {
			route_executors = append(route_executors, pathmethods.Delete.OperationID)
		}
		route_executors = removeDuplicate(route_executors)
	}

	if len(route_executors) > 0 {

		for _, executor := range route_executors {
			Data_schema = Data_schema + fmt.Sprintf("\n    \"%v\":%v,", executor, executor)
			Data_userfunction = Data_userfunction + fmt.Sprintf("\nfunc %v(c *gin.Context) {\n"+
				"    c.JSON(http.StatusOK, gin.H{%v})"+
				"\n}", executor, executor_result[executor])
		}
	}

	// Data_schema = ""
	// "getUsersList":   getUsersList,
	// "getUserDetails": getUserDetails,

	// Data_userfunction = ""
	// func getUserDetails(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"msg": "getUserDetails"})
	// }

	Data_api = ""
	//do nothing at the moment
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
