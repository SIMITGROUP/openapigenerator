package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

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
	}

	var routebytes bytes.Buffer
	routepath := "./templates/go/routehandle.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, allhandles)
	helper.WriteFile("openapi", "routerhandle.go", routebytes.String())
}
