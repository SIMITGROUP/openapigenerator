package buildgo

import (
	"bytes"
	"fmt"
	"openapigenerator/helper"
	"strings"
	"text/template"
)

// prepare unit test request path according path's parameter
func GenerateSamplePath(path string, paras map[string]helper.Model_Parameter) (newpath string) {
	newpath = path
	for paraname, parasetting := range paras {
		if parasetting.StoreIn == "path" {
			replacefrom := fmt.Sprintf(":%v", paraname)
			replaceto := parasetting.Example
			newpath = strings.Replace(newpath, replacefrom, replaceto, -1)
		}
	}
	return
}

// generate meaningful test function name base on request method and path
func GenerateTestFuncName(method string, path string) (funcname string) {
	pathstr := ""
	patharr := strings.Split(path, "/")

	for _, partstr := range patharr {
		partstr = strings.Replace(partstr, "{", "", -1)
		partstr = strings.Replace(partstr, "}", "", -1)
		partstr = strings.Replace(partstr, "-", "", -1)
		partstr = strings.Replace(partstr, ":", "", -1)
		pathstr = pathstr + helper.UpperCaseFirst(partstr)
	}

	funcname = fmt.Sprintf("%v_%v", helper.UpperCaseFirst(method), pathstr)
	return
}
func getTestServer() string {
	applink := "http://localhost"
	arr := strings.Split(helper.Proj.ListenAddress, ":")
	if len(arr) == 2 {
		applink = applink + ":" + arr[1]
	}
	return applink
}
func WriteTest() {
	routesettings := helper.AllRoutes
	for path, pathsetting := range routesettings {
		//convert parameter in path to value, like /user/{user} => /user/myuid
		// newpath := helper.ConvertPathParasCurlyToColon(path)
		// flattenpath := strings.Replace(path, "/", "_", -1)
		// flattenpath = strings.Replace(flattenpath, ":", "_", -1)
		// flattenpath = strings.Replace(flattenpath, "{", "", -1)
		// flattenpath = strings.Replace(flattenpath, "}", "", -1)
		// flattenpath = strings.Replace(flattenpath, "#", "_", -1)
		// flattenpath = strings.Replace(flattenpath, ".", "_", -1)
		// flattenpath = strings.Replace(flattenpath, "-", "_", -1)
		//create test, 1 request method 1 file

		for method, reqsetting := range pathsetting.RequestSettings {
			newpath := GenerateSamplePath(path, reqsetting.RequestHandle.Parameters)
			var writebytes bytes.Buffer
			srcsettings := map[string]string{}
			functionname := GenerateTestFuncName(method, path)

			srcsettings["TestName"] = "Test" + functionname
			srcsettings["RequestServer"] = getTestServer()
			srcsettings["RequestMethod"] = method
			srcsettings["RequestPath"] = newpath
			srcsettings["ContentType"] = reqsetting.RequestHandle.ContentType
			testfilename := functionname + "_test.go"
			srcpath := "./templates/go/test.gotxt"
			src := helper.ReadFile(srcpath)
			srctemplate := template.New("test")
			srctemplate, _ = srctemplate.Parse(src)
			_ = srctemplate.Execute(&writebytes, srcsettings)
			helper.WriteFile("test", testfilename, writebytes.String())
		}

	}
}
