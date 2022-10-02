package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"strings"
	"text/template"
)

// register routes
func WriteRoutes() {
	routesettings := helper.AllRoutes
	for path, pathsetting := range routesettings {
		newpath := helper.ConvertPathParasCurlyToColon(path)
		for method, route := range pathsetting.RequestSettings {
			desc := strings.Replace(route.Description, "\n", "\n    // ", -1)
			route.Description = desc
			route.Path = newpath
			pathsetting.RequestSettings[method] = route
		}

		delete(routesettings, path)
		routesettings[newpath] = pathsetting
	}
	// routesettings["sss"].Path
	// routesettings["sss"].RequestSettings
	var routebytes bytes.Buffer
	routepath := "./templates/go/routeregistry.gotxt"
	routesrc := helper.ReadFile(routepath)
	routetemplate := template.New("route")
	routetemplate, _ = routetemplate.Parse(routesrc)
	_ = routetemplate.Execute(&routebytes, routesettings)
	helper.WriteFile("openapi", "routeregistry.go", routebytes.String())
}
