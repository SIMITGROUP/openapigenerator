package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var AuthMiddleware = map[string]string{}

func prepareSecuritySchemes(schemes openapi3.SecuritySchemes) string {
	securityschemesstr := ""

	for securityname, setting := range schemes {
		AuthMiddleware["securityname"] = ""
		// fmt.Println("securityname", securityname)
		tmp := ""

		switch setting.Value.Type {
		case "apiKey":
			tmp = getAPIKeyAuthStr(securityname, setting.Value)
		case "http":
			if setting.Value.Scheme == "basic" {
				tmp = getBasicHttpAuthStr(securityname, setting.Value)
			} else if setting.Value.Scheme == "bearer" {
				if setting.Value.BearerFormat == "JWT" {
					tmp = getJWTHttpAuthStr(securityname, setting.Value)
				}
			}

			//scheme = basic
			//or
			//scheme =bearer
			//bearerformat: JWT

		//not supported yet
		case "oauth2":
		case "openIdConnect":
		}

		if tmp != "" {
			securityschemesstr = securityschemesstr + "\n" + tmp
		}

		// authname := setting.Value.Name
		// authtype := setting.Value.Type
		// authscheme := setting.Value.Scheme
		// authformat := setting.Value.BearerFormat
		// authin := setting.Value.In
		// authflows := setting.Value.Flows
		// authopenidconnecturl := setting.Value.OpenIdConnectUrl
		// authdesc := setting.Value.Description

		/*
			type: apiKey, http, oauth2, openIdConnect
		*/
	}
	if securityschemesstr != "" {
		securityschemesstr = `package openapi
import (
	"fmt"
	"github.com/gin-gonic/gin"
)` + "\n\n" + securityschemesstr
	}
	writeFile("openapi", "securityschemes.go", securityschemesstr)
	return securityschemesstr
}

func getAPIKeyAuthStr(name string, setting *openapi3.SecurityScheme) string {
	return ""
}

func getBasicHttpAuthStr(name string, setting *openapi3.SecurityScheme) string {
	return ""
}
func getJWTHttpAuthStr(name string, setting *openapi3.SecurityScheme) string {
	template := "type " + name + " struct {token string}\n" +
		"type " + name + "i interface {func" + name + "(c *gin.Context)}\n" +
		"func (obj " + name + ") func" + name + "(c *gin.Context) {fmt.Print(obj)}\n" +
		"var data" + name + " " + name + "i = " + name + "{}\n"
	return template
}
