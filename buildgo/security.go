package buildgo

import (
	"bytes"
	"log"
	"openapigenerator/helper"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

func WriteSecuritySchemes() {

	for authname, setting := range helper.Allsecurityschemas {

		securityscheme := helper.LowerCaseFirst(setting.Value.Scheme)
		securitytype := helper.LowerCaseFirst(setting.Value.Type)
		handlename := getAuthHandles(authname, setting.Value)
		handleData := ""
		if securitytype == "basic" {

		}

		securityobj := helper.Model_Security{
			Name:           authname,
			SecurityType:   securitytype,
			SecurityScheme: securityscheme,
			ModelName:      helper.GetModelName(authname),
			InterfaceName:  helper.GetInterfaceName(authname),
			MethodName:     "Auth_" + authname,
			HandleName:     handlename,
			HandleData:     handleData,
		}

		filename := "Security_" + authname + ".go"

		var writebuffer bytes.Buffer
		schemapath := "./templates/go/security.gotxt"
		schemasrc := helper.ReadFile(schemapath)
		schematemplate := template.New("security")
		schematemplate, _ = schematemplate.Parse(schemasrc)
		// fmt.Println("securityobj::::", securityobj)
		err := schematemplate.Execute(&writebuffer, securityobj)
		if err != nil {
			log.Fatal("writing template ", filename, "error, ", err)
		}
		helper.WriteFile("openapi", filename, writebuffer.String())
		_, _, _, _ = setting, filename, securityobj, writebuffer
	}
}

func getAuthHandles(authname string, setting *openapi3.SecurityScheme) string {
	authstr := "func(c *gin.Context) {}"

	switch setting.Type {
	case "http":
		schemaname := helper.LowerCaseFirst(setting.Scheme)
		if schemaname == "basic" {
			authstr = "gin.BasicAuth(BasicAuthAccounts_" + authname + ")"
		} else if schemaname == "bearer" {
			// authstr = "func(c *gin.Context) {}"
		}
	//not supported yet
	case "apiKey":
	case "mutualTLS":
	case "oauth2":
	case "openIdConnect":
	}
	return authstr
}
