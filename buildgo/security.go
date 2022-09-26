package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/getkin/kin-openapi/openapi3"
)

// consolidate multiple security scheme, every scheme name = separate handle
// different securityscheme type write in different file
// example: split by basichttp, jwt, oauth2
type Model_SchemeHandle struct {
	SchemeType  string
	SchemeName  string
	Auth_handle []string
}

func WriteSecuritySchemes() {

	// fmt.Println("WriteSecuritySchemes")
	schemefilelist := make(map[string]string)
	schemenamelist := make(map[string][]string)

	for authname, setting := range helper.Allsecurityschemas {
		securityscheme := helper.LowerCaseFirst(setting.Value.Scheme)
		securitytype := helper.LowerCaseFirst(setting.Value.Type)
		schemetype, filename, schemehandle := getAuthTemplates(authname, setting.Value)
		schemefilelist[schemetype] = filename
		authhandlename := helper.GetAuthMethodName(authname)
		schemenamelist[schemetype] = append(schemenamelist[schemetype], authhandlename)
		log.Debug("Prepare Scheme: ", authname, ", type: ", schemetype)

		// schemasrc := helper.ReadFile(filename)
		// schematemplate := template.New("security")
		// schematemplate, _ = schematemplate.Parse(schemasrc)
		_, _, _, _, _, _ = setting, filename, securitytype, schemetype, securityscheme, schemehandle
	}

	for schemetype, srcfilename := range schemefilelist {

		filename := "Security_" + schemetype + ".go"
		var writebuffer bytes.Buffer
		schemasrc := helper.ReadFile(srcfilename)
		schematemplate := template.New("security")
		schematemplate, _ = schematemplate.Parse(schemasrc)
		securityobj := Model_SchemeHandle{
			SchemeType:  schemetype,
			SchemeName:  schemetype + "name",
			Auth_handle: schemenamelist[schemetype],
		}
		err := schematemplate.Execute(&writebuffer, securityobj)
		if err != nil {
			log.Fatal("writing template ", filename, "error, ", err)
		}
		helper.WriteFile("openapi", filename, writebuffer.String())
	}
}

func getAuthTemplates(authname string, setting *openapi3.SecurityScheme) (string, string, string) {
	filename := ""
	schemaname := helper.LowerCaseFirst(setting.Scheme)
	schemetype := ""
	schemehandle := ""

	if setting.Type == "http" && schemaname == "basic" {
		schemetype = "basic"
		filename = "security_httpbasic"
		schemehandle = "BasicAuth"
	} else if setting.Type == "http" && schemaname == "bearer" { //JWT
		schemetype = "jwt"
		filename = "security_httpjwt"
		schemehandle = "xxxxxxx"
	} else if setting.Type == "apiKey" {
		schemetype = "apikey"
		filename = "security_apikey"
		schemehandle = "xxxxxxx"
	} else if setting.Type == "mutualTLS" {
	} else if setting.Type == "oauth2" {
	} else if setting.Type == "openIdConnect" {
		schemetype = "openidconnect"
		filename = "security_openidconnect"
		schemehandle = "xxxxxxx"
	}
	if schemetype != "" {
		filename = "templates/go/" + filename + ".gotxt"
		return schemetype, filename, schemehandle
	} else {
		log.Fatal(authname + " using security scheme type " + setting.Type + " is not supported yet")
		return schemetype, filename, schemehandle
	}

}

// func WriteSecuritySchemes2() {

// 	for authname, setting := range helper.Allsecurityschemas {

// 		securityscheme := helper.LowerCaseFirst(setting.Value.Scheme)
// 		securitytype := helper.LowerCaseFirst(setting.Value.Type)
// 		handlename := getAuthHandles(authname, setting.Value)
// 		handleData := ""

// 		securityobj := helper.Model_Security{
// 			Name:           authname,
// 			SecurityType:   securitytype,
// 			SecurityScheme: securityscheme,
// 			ModelName:      helper.GetModelName(authname),
// 			InterfaceName:  helper.GetInterfaceName(authname),
// 			MethodName:     "Auth_" + authname,
// 			HandleName:     handlename,
// 			HandleData:     handleData,
// 		}
// 		fmt.Println("securityobj", securityobj)

// 		filename := "Security_" + authname + ".go"

// 		var writebuffer bytes.Buffer
// 		schemetype, schemapath := getAuthTemplatesFile(authname, setting.Value)

// 		schemasrc := helper.ReadFile(schemapath)
// 		schematemplate := template.New("security")
// 		schematemplate, _ = schematemplate.Parse(schemasrc)
// 		fmt.Println("securityobj::::", securityobj)
// 		err := schematemplate.Execute(&writebuffer, securityobj)
// 		if err != nil {
// 			log.Fatal("writing template ", filename, "error, ", err)
// 		}
// 		helper.WriteFile("openapi", filename, writebuffer.String())
// 		_, _, _, _, _ = setting, filename, securityobj, writebuffer, schemetype
// 	}
// }

// func getAuthHandles(authname string, setting *openapi3.SecurityScheme) string {
// 	authstr := "func(c *gin.Context) {}"

// 	switch setting.Type {

// 	case "http":
// 		schemaname := helper.LowerCaseFirst(setting.Scheme)
// 		if schemaname == "basic" {
// 			authstr = "gin.BasicAuth(BasicAuthAccounts_" + authname + ")"
// 		} else if schemaname == "bearer" {
// 		}
// 	case "apiKey":
// 	case "mutualTLS":
// 	case "oauth2":
// 	case "openIdConnect":
// 	}
// 	return authstr
// }
