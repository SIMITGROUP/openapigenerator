package main

import (
	"flag"

	log "github.com/sirupsen/logrus"

	"openapigenerator/buildgo"
	"openapigenerator/buildphp"
	"openapigenerator/helper"

	"github.com/getkin/kin-openapi/openapi3"
)

var GenerateFolder = ""
var ProjectName = ""
var ApiFile = ""
var ListenAddress = ""
var BuildLang = ""

func main() {
	flag.StringVar(&BuildLang, "lang", "go", "Build language (go/php)")
	flag.StringVar(&GenerateFolder, "targetfolder", "../openapiserverfolder", "Generate Folder to which folder")
	flag.StringVar(&ProjectName, "projectname", "openapiserver", "Rest API GO project name")
	flag.StringVar(&ApiFile, "apifile", "spec.yaml", "openapi v3 yaml file")
	flag.StringVar(&ListenAddress, "listen", ":8982", "listen to which address, default :8982")
	flag.Parse()
	helper.Proj.ApiFile = ApiFile
	helper.Proj.ListenAddress = ListenAddress
	helper.Proj.BuildLang = BuildLang
	helper.Proj.GenerateFolder = GenerateFolder
	helper.Proj.ProjectName = ProjectName
	log.SetLevel(log.DebugLevel)
	switch BuildLang {
	case "go":
		GenerateCode(ApiFile)
	default:
		log.Fatal("Build '" + BuildLang + "' is not supported, use 'go' instead!")
	}

}

func GenerateCode(ApiFile string) {
	doc, err := openapi3.NewLoader().LoadFromFile(ApiFile)
	if err != nil {
		log.Fatal(err)
	}
	helper.PrepareObjects(doc)
	/*






	 */
	// helper.ReadRoutes(doc)
	// helper.ReadComponents(doc)
	switch BuildLang {
	case "go":
		buildgo.Generate(doc)
	case "php":
		buildphp.Generate(doc)
	default:
		log.Fatal("only 'go' is supported at the moment")
		// buildphp.Generate("")
	}

}

// func ReadSchema() {

// }

// func PrepareInfra(doc *openapi3.T) {
// 	switch BuildLang {
// 	case "go":
// 		buildgo.WriteInfra() //Generate(doc)
// 		buildgo.PrepareRoutes(doc)
// 		buildgo.PrepareRouteHandles(doc)

// 	case "php":
// 		// buildphp.PrepareInfra() //Generate(doc)
// 	// case "typescript":
// 	// case "java":
// 	// 	buildphp.Generate(doc)
// 	default:
// 		fmt.Println(BuildLang + " build is not supported")
// 	}
// }
