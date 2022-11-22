package buildphp

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

func Generate(doc *openapi3.T) {
	log.Info("Generate PHP API:")
	WriteInfra()
	// WriteSchemas()
	// WriteRoutes()
	// WriteSecuritySchemes()
	// WriteHandles()
	// WriteTest()
}

func WriteInfra() {
	// prepare main
	var mainbytes bytes.Buffer
	mainfilepath := "templates/php/index.phptxt"
	mainsrc := helper.ReadFile(mainfilepath)
	maintemplate := template.New("main")
	maintemplate, _ = maintemplate.Parse(mainsrc)
	_ = maintemplate.Execute(&mainbytes, helper.Proj)
	mainstr := mainbytes.String()
	// mainstr = template.UnescapeString(mainstr)
	helper.WriteFile("public", "index.php", mainstr)

	// prepare coposer
	var composerbytes bytes.Buffer
	composerfilepath := "templates/php/composer.json"
	composersrc := helper.ReadFile(composerfilepath)
	composertemplate := template.New("composer")
	composertemplate, _ = composertemplate.Parse(composersrc)
	_ = composertemplate.Execute(&composerbytes, helper.Proj)
	helper.WriteFile("", "composer.json", composerbytes.String())

	// // prepare Makefile
	// var makebytes bytes.Buffer
	// makefilepath := "templates/go/Makefile.txt"
	// makesrc := helper.ReadFile(makefilepath)
	// maketemplate := template.New("makefile")
	// maketemplate, _ = maketemplate.Parse(makesrc)
	// _ = maketemplate.Execute(&makebytes, helper.Proj)
	// helper.WriteFile("", "Makefile", makebytes.String())

	// //prepare openapi/server.go

	// var serverbytes bytes.Buffer
	// serverfilepath := "templates/go/server.gotxt"
	// serversrc := helper.ReadFile(serverfilepath)
	// servertemplate := template.New("server")
	// servertemplate, _ = servertemplate.Parse(serversrc)
	// _ = servertemplate.Execute(&serverbytes, helper.Proj)
	// helper.WriteFile("openapi", "server.go", serverbytes.String())
}
