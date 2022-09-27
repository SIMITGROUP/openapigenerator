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
type Model_SchemeHandles struct {
	SchemeType  string
	SchemeName  string
	Auth_handle []Model_SchemeSetting
}
type Model_SchemeSetting struct {
	Keyname     string
	Handlename  string
	Description string
}

func WriteSecuritySchemes() {

	// fmt.Println("WriteSecuritySchemes")
	schemefilelist := make(map[string]string)
	schemenamelist := make(map[string][]Model_SchemeSetting)

	for authname, setting := range helper.Allsecurityschemas {
		securityscheme := helper.LowerCaseFirst(setting.Value.Scheme)
		securitytype := helper.LowerCaseFirst(setting.Value.Type)
		schemetype, filename, schemehandle := getAuthTemplates(authname, setting.Value)
		schemefilelist[schemetype] = filename
		authhandlename := helper.GetAuthMethodName(authname)
		schemenamelist[schemetype] = append(schemenamelist[schemetype],
			Model_SchemeSetting{
				Handlename:  authhandlename,
				Keyname:     setting.Value.Name,
				Description: setting.Value.Description,
			})
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
		securityobj := Model_SchemeHandles{
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
	} else if setting.Type == "apiKey" {
		// setting
		schemetype = "apikey"
		filename = "security_apikey"
		schemehandle = "xxxxxxx"
		keyname := setting.Name
		if verifyKeyname(keyname) == false {
			log.Fatal("Invalid apikey " + keyname + ", it should only consist character a-z without special character and spacing")
		}
		keyin := setting.In
		if keyin != "header" {
			log.Fatal("Api key shall define 'In' value as header")
		}
		keydesc := setting.Description
		_ = keydesc
	} else if setting.Type == "http" && schemaname == "bearer" { //JWT
		log.Fatal("JWT security scheme is not supported yet")
		schemetype = "jwt"
		filename = "security_httpjwt"
		schemehandle = "xxxxxxx"
	} else if setting.Type == "mutualTLS" {
		log.Fatal("mutualTLS security scheme is not supported yet")
	} else if setting.Type == "oauth2" {
		log.Fatal("oauth2 security scheme is not supported yet")
	} else if setting.Type == "openIdConnect" {
		log.Fatal("openIdConnect security scheme is not supported yet")
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
func verifyKeyname(s string) bool {
	for _, r := range s {
		if r == '_' {
			return true
		} else if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
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
