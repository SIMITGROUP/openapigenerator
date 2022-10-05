package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"strings"
	"text/template"
)

// write unit test of each route
func WriteTest() {
	routesettings := helper.AllRoutes
	for path, pathsetting := range routesettings {
		//convert parameter in path to value, like /user/{user} => /user/myuid
		newpath := helper.ConvertPathParasCurlyToColon(path)
		flattenpath := strings.Replace(path, "/", "_", -1)
		flattenpath = strings.Replace(flattenpath, "{", "", -1)
		flattenpath = strings.Replace(flattenpath, "}", "", -1)
		flattenpath = strings.Replace(flattenpath, "#", "_", -1)
		flattenpath = strings.Replace(flattenpath, ".", "_", -1)
		flattenpath = strings.Replace(flattenpath, "-", "_", -1)
		//create test, 1 request method 1 file
		for method, _ := range pathsetting.RequestSettings {

			var writebytes bytes.Buffer
			srcsettings := map[string]string{}
			srcsettings["TestName"] = method + "_" + flattenpath
			srcsettings["RequestServer"] = "http://localhost:8000"
			srcsettings["RequestMethod"] = method
			srcsettings["RequestPath"] = newpath

			testfilename := helper.LowerCaseFirst(method) + "_" + flattenpath + "_test.go"
			srcpath := "./templates/go/test.gotxt"
			src := helper.ReadFile(srcpath)
			srctemplate := template.New("test")
			srctemplate, _ = srctemplate.Parse(src)
			_ = srctemplate.Execute(&writebytes, srcsettings)
			helper.WriteFile("test", testfilename, writebytes.String())
		}
		_ = newpath
	}
}
