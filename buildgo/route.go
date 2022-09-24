package buildgo

import (
	"bytes"
	"fmt"
	"openapigenerator/helper"
	"strings"
	"text/template"
)

// register routes
func WriteRoutes() {
	routesettings := helper.GetRoutes()
	var routebytes bytes.Buffer
	routepath := "./templates/go/routeregistry.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, routesettings)
	helper.WriteFile("openapi", "routeregistry.go", routebytes.String())
}

// use every registered route's operationID to create handle function
func WriteHandles() {

	//prepare unique route handles
	alloperations := make(map[string]string)

	for _, route := range helper.Allroutes {
		// fmt.Println(method, path, content.Schema.Ref, datatype)
		if route.DataType != "" {
			alloperations[route.OperationID] = route.DataType
		} else {
			alloperations[route.OperationID] = "gin.H"
		}
	}

	//prepare object for draw template
	handlelist := []helper.Model_Handle{}
	templateobj := helper.Model_HandleTemplate{}

	for handlename, handledatatype := range alloperations {
		// fmt.Println("handlename", handlename, "Datatype", handledatatype)
		oridatatype := strings.Replace(handledatatype, "Model_", "", -1)
		schemobj := helper.GetSchemaFromName(oridatatype)

		if schemobj.Type == "array" {

			handledatatype = "[]" + helper.GetModelNameFromRef(schemobj.Items.Ref)
		}
		fmt.Println("handle info", handlename, handledatatype)
		h := helper.Model_Handle{
			FuncName:   handlename,
			DataType:   handledatatype,
			SchemaType: schemobj.Type,
		}
		handlelist = append(handlelist, h)
	}
	templateobj.Handles = handlelist
	// for _, route := range routesettings.Methods {

	// }
	// x := helper.MethodSettings{}
	// x.Description
	// x.Middlewares
	// x.Method
	// x.OperationID
	// x.Summary
	// x.Path
	// fmt.Println("templateobj", templateobj)
	var routebytes bytes.Buffer
	routepath := "./templates/go/routehandle.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, templateobj)
	helper.WriteFile("openapi", "routerhandle.go", routebytes.String())

}

// map route to opeartion function
// func PrepareRouteHandles(doc *openapi3.T) {

// 	//get unique operationID
// 	var list []helper.Model_Handle
// 	alloperations := []string{}
// 	for _, route := range Allroutes {
// 		alloperations = append(alloperations, route.OperationID)
// 	}
// 	alloperations = helper.RemoveDuplicate(alloperations)

// 	for _, handlename := range alloperations {
// 		list = append(list, helper.Model_Handle{handlename, "gin.H", ""})
// 	}
// 	fmt.Println(list)
// 	handlesettings := helper.HandleTemplateObj{Handles: list}
// 	var handlebytes bytes.Buffer
// 	handlepath := "./templates/go/routehandle.gotxt"
// 	handlesrc := helper.ReadFile(handlepath)
// 	handletemplate := template.New("handles")
// 	handletemplate, _ = handletemplate.Parse(handlesrc)
// 	_ = handletemplate.Execute(&handlebytes, handlesettings)
// 	helper.WriteFile("openapi", "routehandle.go", handlebytes.String())

// 	// PrepareResponses()
// }
