package buildgo

import (
	"bytes"
	"html/template"
	"openapigenerator/helper"

	log "github.com/sirupsen/logrus"
)

// use every registered route's operationID to create handle function
func WriteHandles() {
	allhandles := map[string]helper.Model_RequestHandle{}
	for handlename, hsetting := range helper.AllRequestHandles {
		// fmt.Println("handlename", handlename, "Datatype", handledatatype)

		if hsetting.ResponseSchema.ModelType == "array" {
			// handledatatype = "[]" + helper.GetModelNameFromRef(schemobj.Items.Ref)
		}

		log.Info("handle info", handlename, ", ", hsetting.ResponseSchema.ModelType)
		h := helper.Model_RequestHandle{
			HandleName:     handlename,
			ResponseSchema: hsetting.ResponseSchema,
			Parameters:     hsetting.Parameters,
			RequestBodies:  hsetting.RequestBodies,
		}
		allhandles[handlename] = h
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

	}
	var routebytes bytes.Buffer
	routepath := "./templates/go/routehandle.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, allhandles)
	helper.WriteFile("openapi", "routerhandle.go", routebytes.String())
}
