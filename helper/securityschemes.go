package helper

import (
	log "github.com/sirupsen/logrus"
)

func PrepareSecuritySchemes() {
	securityschemas := Doc.Components.SecuritySchemes
	var scopes map[string]string
	for authname, setting := range securityschemas {
		log.Info("Preparing Security Scheme: ", authname)
		sc := *setting.Value
		schemetype := LowerCaseFirst(setting.Value.Type)
		// AllFunctionName = append(AllFunctionName, GetAuthMethodName(authname))
		if schemetype == "oauth2" {
			Proj.InitFunctions = append(Proj.InitFunctions, authname+"_prepare")

			Proj.AllEnvVars[authname+"_configURL"] = ""
			Proj.AllEnvVars[authname+"_redirectURL"] = ""
			Proj.AllEnvVars[authname+"_clientID"] = ""
			Proj.AllEnvVars[authname+"_clientSecret"] = ""

			if sc.Flows.Password != nil {
				scopes = sc.Flows.Password.Scopes
			} else if sc.Flows.Implicit != nil {
				scopes = sc.Flows.Implicit.Scopes
			} else if sc.Flows.ClientCredentials != nil {
				scopes = sc.Flows.ClientCredentials.Scopes
			}
		} else if schemetype == "apikey" {
			envkeyname := ConvertApiKeyToEnvVar(setting.Value.Name)
			Proj.AllEnvVars[envkeyname] = ""
		}

		AllSecuritySchemes[authname] = Model_SecuritySchemaSetting{
			Type:        sc.Type,
			Description: sc.Description,
			In:          sc.In,
			SchemeName:  authname,
			Name:        sc.Name,
			Scheme:      sc.Scheme,
			Scopes:      scopes,
		}

		log.Info("Prepare Security Scheme: ", authname, ", type: ", schemetype)

	}
}
