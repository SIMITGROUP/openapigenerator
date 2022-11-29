package buildgo

import (
	"bytes"
	"fmt"
	"openapigenerator/helper"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
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
	// arr := strings.Split(helper.Proj.ListenAddress, ":")
	// if len(arr) == 2 {
	applink = applink + ":" + helper.Proj.ListenPort
	// }
	return applink
}
func WriteTest() {
	routesettings := helper.AllRoutes
	for path, pathsetting := range routesettings {
		for method, reqsetting := range pathsetting.RequestSettings {
			newpath := GenerateSamplePath(path, reqsetting.RequestHandle.Parameters)
			var writebytes bytes.Buffer
			srcsettings := map[string]any{}
			functionname := GenerateTestFuncName(method, path)

			srcsettings["FuncName"] = functionname
			srcsettings["RequestServer"] = getTestServer()
			srcsettings["RequestMethod"] = method
			srcsettings["RequestPath"] = newpath
			srcsettings["ContentType"] = reqsetting.RequestHandle.ContentType
			srcsettings["StatusCode"] = strconv.FormatInt(int64(reqsetting.RequestHandle.HttpStatusCode), 10)
			srcsettings["Envvars"] = helper.Proj.AllEnvVars
			//if requestbody is required, prepare sample requestbody
			if reqsetting.RequestHandle.RequestBodies.RequestSchema.ModelType != "" {
				srcsettings["WithRequestBody"] = true

				modelname := reqsetting.RequestHandle.RequestBodies.RequestSchema.ModelName
				var schemaobj helper.Model_SchemaSetting
				log.Error("functionname:", functionname, "  ,modelname:", modelname)
				for schemaname, schm := range helper.AllSchemas {
					if schm.ModelName == modelname {
						log.Error(" .     schema name:", schemaname)
						schemaobj = helper.AllSchemas[schemaname]
						break
					}
				}

				examplesstr := "`{\n"
				for f, fsetting := range schemaobj.FieldList {
					tmp := fmt.Sprintf("\"%s\": %s,\n", f, fsetting.Example)
					tmp = strings.Replace(tmp, "map[string]interface {}", "", -1)
					examplesstr = examplesstr + tmp

				}
				examplesstr = strings.TrimSuffix(examplesstr, ",\n")
				examplesstr = examplesstr + "}`"
				log.Error("examplesstr:", examplesstr)
				srcsettings["Examplestr"] = examplesstr
			} else {
				srcsettings["Examplestr"] = ""
			}

			testfilename := "Z" + functionname + "_test.go"
			srcpath := "templates/go/test.gotxt"
			src := helper.ReadFile(srcpath)
			srctemplate := template.New("test")
			srctemplate, _ = srctemplate.Parse(src)
			_ = srctemplate.Execute(&writebytes, srcsettings)
			helper.WriteFile("test", testfilename, writebytes.String())
		}

	}
}
