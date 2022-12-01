package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// use every registered route's operationID to create handle function
func WriteHandles() bool {
	// overridefile := helper.Proj.OverrideHandle

	// //if file exists, but no override flag
	// if helper.CheckFileExists("openapi", "routerhandle.go") && overridefile == false {
	// 	log.Warn("routerhandle.go exists, skip")
	// 	return false
	// }

	allhandles := map[string]helper.Model_RequestHandle{}
	for handlename, hsetting := range helper.AllRequestHandles {
		// fmt.Println("handlename", handlename, "Datatype", handledatatype)

		if hsetting.ResponseSchema.ModelType == "array" {
			// handledatatype = "[]" + helper.GetModelNameFromRef(schemobj.Items.Ref)
		}
		firstcharacter := helper.Left(handlename, 1)
		if firstcharacter == "-" {

			log.Warn("Skip generate ", handlename)
			delete(allhandles, handlename)
		} else {
			log.Info("handle: ", handlename, ", ", hsetting.HttpStatusCode, " ", hsetting.ContentType)
			existvalue, exists := helper.Proj.AllExistsHandles[handlename]
			//no generate route handle, if api document exclude it
			if exists == false || existvalue == false {
				allhandles[handlename] = hsetting
			}

		}

	}

	var routebytes bytes.Buffer

	routepath := "templates/go/routehandle.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, allhandles)
	helper.WriteFile("openapi", "ZRouterHandle.go", routebytes.String())
	return true
}
