package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"strings"
	"text/template"

	"gopkg.in/mgo.v2/bson"
)

// register routes
func WriteRoutes() {
	routesettings := helper.AllRoutes
	for path, pathsetting := range routesettings {
		newpath := helper.ConvertPathParasCurlyToColon(path)
		for method, route := range pathsetting.RequestSettings {
			route.Description = strings.Replace(route.Description, "\n", "\n    // ", -1)
			route.Path = newpath

			route.RequestHandle.HandleName = strings.Replace(route.RequestHandle.HandleName, "-", "", -1)
			pathsetting.RequestSettings[method] = route
		}

		delete(routesettings, path)
		routesettings[newpath] = pathsetting
	}
	// routesettings["sss"].Path
	// routesettings["sss"].RequestSettings

	parameters := bson.M{
		"routesettings":  routesettings,
		"projectsetting": helper.Proj,
	}
	var routebytes bytes.Buffer
	routepath := "templates/go/routeregistry.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	// _ = routetemplate.Execute(&routebytes, routesettings)
	_ = routetemplate.Execute(&routebytes, parameters)
	helper.WriteFile("openapi", "ZRouterRegistry.go", routebytes.String())
}
