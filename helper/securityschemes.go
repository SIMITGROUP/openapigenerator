package helper

import (
	log "github.com/sirupsen/logrus"
)

func PrepareSecuritySchemes() {
	securityschemas := Doc.Components.SecuritySchemes
	for authname, setting := range securityschemas {
		schemetype := setting.Value.Type
		AllSecuritySchemes[authname] = *setting.Value
		log.Info("Prepare Security Scheme: ", authname, ", type: ", schemetype)

	}
}
