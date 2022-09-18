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

		schema := doc.Components.Schemas
		schemastring := prepareSchema(schema)
		schemafile := "openapi/schema.go"
		interfacefile := "openapi/interface.go"
		openapifile := "openapi/openapi.go"
		userfunctionfile := "openapi/userfunction.go"

		_ = os.WriteFile(interfacefile, prepareInterface(), 0644)
		_ = os.WriteFile(openapifile, prepareApi(), 0644)
		_ = os.WriteFile(userfunctionfile, prepareUserFunction(), 0644)
		_ = os.WriteFile(schemafile, schemastring, 0644)

		fmt.Println("microservices code generated ", interfacefile, openapifile, userfunctionfile, schemafile)
		fmt.Printf("Edit dummy functions in %v to produce real api result", userfunctionfile)

	}
}

func prepareInterface() []byte {
	data := strings.Replace(Temp_interface, "##data##", Data_interface, -1)
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
func getFieldSettingStr(name string, s openapi3.Schema) string {
	fieldtype := s.Type
	prefix := "    "
	if s.Type == "integer" {
		if s.Format != "" {
			fieldtype = s.Format
		} else {
			fieldtype = "int32"
		}

	} else if s.Type == "object" {
		// fieldtype = " string //original is object"
		tmp := ""
		for subfieldname, subfieldsetting := range s.Properties {
			tmp = tmp + prefix + getFieldSettingStr(subfieldname, *subfieldsetting.Value) + "\n"
		}
		return prefix + name + " struct{\n" + tmp + prefix + "}"
	}
	return prefix + name + " " + fieldtype
}
func prepareSchema(schemas openapi3.Schemas) []byte {

	schemastr := ""
	for schemaname, setting := range schemas {
		// fmt.Println("schema:", schemaname, setting.Value.Title)
		props := setting.Value.Properties
		tmp := ""
		for field, fieldsetting := range props {
			tmp = tmp + getFieldSettingStr(field, *fieldsetting.Value) + "\n"
		}
		schemastr = schemastr + "\ntype " + schemaname + " struct {\n" + tmp + "}\n"
	}
	Data_schema = schemastr
	data := strings.Replace(Temp_schema, "##data##", Data_schema, -1)
	return []byte(data)
}

func getResponses(res openapi3.Responses) string {
	result := ""
	for _, setting := range res {
		content := setting.Value.Content["application/json"]

		// only return first 1, usually http status 200
		if content != nil {
			values := strings.Split(content.Schema.Ref, "/")
			result = values[len(values)-1]
			break
		}
	}
	return result

}
func getFieldTypeSettings(setting *openapi3.Schema) (string, string) {
	fieldtype := setting.Type
	fieldformat := setting.Format
	return fieldtype, fieldformat
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
			Data_interface = Data_interface + fmt.Sprintf("\n    \"%v\":%v,", executor, executor)

			Data_userfunction = Data_userfunction +
				fmt.Sprintf("\nfunc %v(c *gin.Context) {\n"+
					"    data := %v{}\n"+
					"    c.JSON(http.StatusOK, data)"+
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
