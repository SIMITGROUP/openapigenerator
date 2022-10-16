package helper

import (
	log "github.com/sirupsen/logrus"
)

func PrepareSecuritySchemes() {
	securityschemas := Doc.Components.SecuritySchemes
	for authname, setting := range securityschemas {
		schemetype := LowerCaseFirst(setting.Value.Type)
		AllFunctionName = append(AllFunctionName, GetAuthMethodName(authname))
		if schemetype == "oauth2" {
			Proj.InitFunctions = append(Proj.InitFunctions, authname+"_prepare")
			AllFunctionName = append(AllFunctionName, authname+"_login")
			AllFunctionName = append(AllFunctionName, authname+"_logout")
			AllFunctionName = append(AllFunctionName, authname+"_callback")
			AllFunctionName = append(AllFunctionName, authname+"_prepare")

		}
		AllSecuritySchemes[authname] = *setting.Value
		log.Info("Prepare Security Scheme: ", authname, ", type: ", schemetype)

	}
}
