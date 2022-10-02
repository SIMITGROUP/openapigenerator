package buildgo

import (
	"bytes"
	"html/template"
	"openapigenerator/helper"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

// register securityschemes
func WriteSecuritySchemes() {
	log.Info("Write Security Schemes:")
	log.Debug(helper.AllSecuritySchemes)
	for authname, authsetting := range helper.AllSecuritySchemes {
		log.Info("    ", authname)
		data := map[string]string{}
		data["SchemeName"] = authname

		var writebytes bytes.Buffer
		path := getTemplateFile(authsetting)
		targetfile := "Security_" + authname + ".go"
		log.Debug("template source:", path)
		src := helper.ReadFile(path)

		routetemplate := template.New("route")
		routetemplate, _ = routetemplate.Parse(src)
		_ = routetemplate.Execute(&writebytes, data)
		helper.WriteFile("openapi", targetfile, writebytes.String())
	}
}

func getTemplateFile(setting openapi3.SecurityScheme) string {
	filename := ""
	schemaname := helper.LowerCaseFirst(setting.Scheme)

	if setting.Type == "http" && schemaname == "basic" {
		filename = "security_httpbasic"

	} else if setting.Type == "apiKey" {
		// setting
		filename = "security_apikey"
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
		filename = "security_httpjwt"
	} else if setting.Type == "mutualTLS" {
		log.Fatal("mutualTLS security scheme is not supported yet")
	} else if setting.Type == "oauth2" {
		log.Fatal("oauth2 security scheme is not supported yet")
	} else if setting.Type == "openIdConnect" {
		log.Fatal("openIdConnect security scheme is not supported yet")
		filename = "security_openidconnect"
	}

	filename = "templates/go/" + filename + ".gotxt"
	return filename

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
