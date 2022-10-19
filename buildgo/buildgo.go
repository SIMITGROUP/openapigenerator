package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

func Generate(doc *openapi3.T) {
	log.Info("Generate Golang API:")
	WriteInfra()
	WriteSchemas()
	WriteRoutes()
	WriteSecuritySchemes()
	WriteHandles()
	WriteTest()

}

func WriteInfra() {
	overridefile := helper.Proj.OverrideHandle
	//prepare main
	//if file exists, but no override flag
	if helper.CheckFileExists("", "main.go") && overridefile == false {
		log.Warn("main.go exists, skip")
	} else {
		var mainbytes bytes.Buffer
		mainfilepath := "./templates/go/main.gotxt"
		mainsrc := helper.ReadFile(mainfilepath)
		maintemplate := template.New("main")
		maintemplate, _ = maintemplate.Parse(mainsrc)
		_ = maintemplate.Execute(&mainbytes, helper.Proj)
		helper.WriteFile("", "main.go", mainbytes.String())
	}

	//prepare go.mod
	var gomodbytes bytes.Buffer
	gomodfilepath := "./templates/go/go.modtxt"
	gosrc := helper.ReadFile(gomodfilepath)
	gotemplate := template.New("gomod")
	gotemplate, _ = gotemplate.Parse(gosrc)
	_ = gotemplate.Execute(&gomodbytes, helper.Proj)
	helper.WriteFile("", "go.mod", gomodbytes.String())

	//prepare Makefile
	var makebytes bytes.Buffer
	makefilepath := "./templates/go/Makefile.txt"
	makesrc := helper.ReadFile(makefilepath)
	maketemplate := template.New("makefile")
	maketemplate, _ = maketemplate.Parse(makesrc)
	_ = maketemplate.Execute(&makebytes, helper.Proj)
	helper.WriteFile("", "Makefile", makebytes.String())

	//prepare openapi/server.go

	var serverbytes bytes.Buffer
	serverfilepath := "./templates/go/server.gotxt"
	serversrc := helper.ReadFile(serverfilepath)
	servertemplate := template.New("server")
	servertemplate, _ = servertemplate.Parse(serversrc)
	_ = servertemplate.Execute(&serverbytes, helper.Proj)
	helper.WriteFile("openapi", "ZServer.go", serverbytes.String())

	//prepare .env
	var envvarbytes bytes.Buffer
	envvarfilepath := "./templates/go/dot.env"
	dotenvsrc := helper.ReadFile(envvarfilepath)
	envvartemplate := template.New("envvar")
	envvartemplate, _ = envvartemplate.Parse(dotenvsrc)
	_ = envvartemplate.Execute(&envvarbytes, helper.Proj.AllEnvVars)
	helper.WriteFile("", ".env.default", envvarbytes.String())

}
