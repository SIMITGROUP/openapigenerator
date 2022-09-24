package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
)

func Generate(doc *openapi3.T) {
	WriteInfra()
	WriteSchemas()
	WriteRoutes()
	WriteHandles()
	WriteSecuritySchemes()
	// prepareComponent(doc)

	// prepareHandles(doc)
}

func WriteInfra() {

	//prepare main
	var mainbytes bytes.Buffer
	mainfilepath := "./templates/go/main.gotxt"
	mainsrc := helper.ReadFile(mainfilepath)
	maintemplate := template.New("main")
	maintemplate, _ = maintemplate.Parse(mainsrc)
	_ = maintemplate.Execute(&mainbytes, helper.Proj)
	helper.WriteFile("", "main.go", mainbytes.String())

	//prepare go.mod
	var gomodbytes bytes.Buffer
	gomodfilepath := "./templates/go/go.modtxt"
	gosrc := helper.ReadFile(gomodfilepath)
	gotemplate := template.New("gomod")
	gotemplate, _ = gotemplate.Parse(gosrc)
	_ = gotemplate.Execute(&gomodbytes, helper.Proj)
	helper.WriteFile("", "go.mod", gomodbytes.String())

	//prepare openapi/server.go
	var serverbytes bytes.Buffer
	serverfilepath := "./templates/go/server.gotxt"
	serversrc := helper.ReadFile(serverfilepath)
	servertemplate := template.New("server")
	servertemplate, _ = servertemplate.Parse(serversrc)
	_ = servertemplate.Execute(&serverbytes, helper.Proj)
	helper.WriteFile("openapi", "server.go", serverbytes.String())
}
