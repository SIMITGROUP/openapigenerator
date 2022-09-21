package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func preparePaths(doc *openapi3.T) {
	routestr := ""

	for oripath, pathmethods := range doc.Paths {
		path := convertGinPath(oripath)

		if pathmethods.Get != nil {
			routestr = routestr + getRouteString("GET", path, pathmethods.Get)
		}
		if pathmethods.Post != nil {
			routestr = routestr + getRouteString("POST", path, pathmethods.Post)
		}
		if pathmethods.Put != nil {
			routestr = routestr + getRouteString("PUT", path, pathmethods.Put)
		}
		if pathmethods.Delete != nil {
			routestr = routestr + getRouteString("DELETE", path, pathmethods.Delete)
		}
		if pathmethods.Head != nil {
			routestr = routestr + getRouteString("Head", path, pathmethods.Head)
		}
		if pathmethods.Patch != nil {
			routestr = routestr + getRouteString("Patch", path, pathmethods.Patch)
		}
		if pathmethods.Options != nil {
			routestr = routestr + getRouteString("OPTIONS", path, pathmethods.Options)
		}
		if pathmethods.Trace != nil {
			routestr = routestr + getRouteString("TRACE", path, pathmethods.Trace)
		}
	}
	// fmt.Println(routestr)
	filename := "route.go"
	template := "package openapi\n\n" +
		"import \"github.com/gin-gonic/gin\"\n\n" +
		"func addRoute(r *gin.Engine) {%v\n}"

	content := fmt.Sprintf(template, routestr)
	// _ = os.Remove(filename)
	// _ = os.WriteFile(filename, []byte(content), 0644)
	writeFile("openapi", filename, content)

}

func getRouteString(httpmethod string, path string, op *openapi3.Operation) string {
	operatingID := op.OperationID
	securities := op.Security
	routestr := ""
	if securities == nil {
		routestr = fmt.Sprintf("\n    r.%v(\"%v\", %v)", httpmethod, path, operatingID)
	} else {
		handlestring := ""
		for _, securitysetting := range *securities {
			for authname, authsetting := range securitysetting {
				handlestring = handlestring + "data" + authname + ".func" + authname + ", "
				// fmt.Println("auth", authname, authsetting)
				_ = authsetting
			}
		}
		handlestring = handlestring + operatingID
		descriptions := strings.Replace(op.Description, "\n", "//\n", -1)
		templatestr := `
    // %v
    // %v
    r.%v("%v", %v)`
		routestr = fmt.Sprintf(templatestr, op.Summary, descriptions, httpmethod, path, handlestring)
	}

	return routestr
}

/**************************** tools **************************/
func convertGinPath(oripath string) string {
	newpath := oripath
	r := regexp.MustCompile(`{\s*(.*?)\s*}`)
	matches := r.FindAllStringSubmatch(newpath, -1)
	for _, v := range matches {
		openapistr := "{" + v[1] + "}"
		reststr := ":" + v[1]
		newpath = strings.Replace(newpath, openapistr, reststr, -1)
	}

	return newpath
}
