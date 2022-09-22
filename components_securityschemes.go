package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var AuthMiddleware = map[string]string{}

func prepareSecuritySchemes(schemes openapi3.SecuritySchemes) string {
	securityschemesstr := ""
	otherschemestr := ""
	for securityname, setting := range schemes {
		AuthMiddleware["securityname"] = ""
		handlestr, othersstr := getAuthStr(securityname, setting.Value)

		if handlestr != "" {
			securityschemesstr = securityschemesstr + "\n" + handlestr
		}
		if othersstr != "" {
			otherschemestr = otherschemestr + "\n" + othersstr
		}
	}
	if securityschemesstr != "" {
		securityschemesstr = `package openapi
import (
	//"net/http"
	"github.com/gin-gonic/gin"
)` + "\n\n" + securityschemesstr + "\n\n" + otherschemestr
	}
	writeFile("openapi", "securityschemes.go", securityschemesstr)
	return securityschemesstr
}

func getAuthStr(name string, setting *openapi3.SecurityScheme) (string, string) {
	modelname := GetModelName(name)
	interfacename := GetInterfaceName(name)
	methodname := GetAuthMethodName(name)

	template := "type " + modelname + " struct {token string}\n" +
		"type " + interfacename + "i interface {" + methodname + "() gin.HandlerFunc }\n" +
		"func (obj " + modelname + ") " + methodname + "() gin.HandlerFunc {" +
		`return ` + getAuthHandles(setting) +
		"}\n" +
		"var " + name + " " + interfacename + "i = " + modelname + "{}\n"
	supportstr := getSupportString(setting)
	return template, supportstr
}
func getSupportString(setting *openapi3.SecurityScheme) string {
	supportstr := ""
	switch setting.Type {
	case "http":
		schemaname := lowerCaseFirst(setting.Scheme)
		if schemaname == "basic" {
			supportstr = `
//change below for basic authentication users/password
var BasicAuthAccounts = gin.Accounts{
	"admin":    "admin",
	"test": "test",	
}
`
		}
	}
	return supportstr
}
func getAuthHandles(setting *openapi3.SecurityScheme) string {
	authstr := "func(c *gin.Context) {}"

	switch setting.Type {
	case "http":
		schemaname := lowerCaseFirst(setting.Scheme)
		if schemaname == "basic" {
			authstr = "gin.BasicAuth(BasicAuthAccounts)"
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

// func getAPIKeyAuthStr(name string, setting *openapi3.SecurityScheme) string {
// 	return ""
// }

// func getBasicHttpAuthStr(name string, setting *openapi3.SecurityScheme) string {
// 	modelname := GetModelName(name)
// 	interfacename := GetInterfaceName(name)
// 	methodname := GetAuthMethodName(name)

// 	template := "type " + modelname + " struct {token string}\n" +
// 		"type " + interfacename + "i interface {" + methodname + "() gin.HandlerFunc }\n" +
// 		"func (obj " + modelname + ") " + methodname + "() gin.HandlerFunc {" +
// 		`return gin.BasicAuth(gin.Accounts{})` +
// 		"}\n" +
// 		"var " + name + " " + interfacename + "i = " + modelname + "{}\n"
// 	return template
// }
// func getJWTHttpAuthStr(name string, setting *openapi3.SecurityScheme) string {
// 	modelname := GetModelName(name)
// 	interfacename := GetInterfaceName(name)
// 	methodname := GetAuthMethodName(name)
// 	template := "type " + modelname + " struct {token string}\n" +
// 		"type " + interfacename + "i interface {" + methodname + "() gin.HandlerFunc }\n" +
// 		"func (obj " + modelname + ") " + methodname + "()  gin.HandlerFunc {" +
// 		`return func(c *gin.Context) {
// 			data := gin.H{
// 				"msg": "no authenticated",
// 			}
// 			c.JSON(http.StatusBadRequest, data)
// 			c.Abort()
// 		}` +
// 		"}\n" +
// 		"var " + name + " " + interfacename + "i = " + modelname + "{}\n"
// 	return template
// }
