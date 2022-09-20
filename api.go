package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

var mainsource = `package main

import "%v/openapi"

func main() {
	r := openapi.NewServer()
	r.Run(":%v")	
}

`
var gomodulesource = `module %v
go 1.19`

var apisource = `package openapi

import (
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	addRoute(r)
	return r
}

`

func Generate(docfile string) {
	writeFile("", "main.go", fmt.Sprintf(mainsource, ProjectName, Port))
	writeFile("", "go.mod", fmt.Sprintf(gomodulesource, ProjectName))
	writeFile("openapi", "openapi.go", apisource)
	doc, _ := openapi3.NewLoader().LoadFromFile(docfile)
	prepareComponent(doc)
	preparePaths(doc)
	prepareHandles(doc)

	// cli := "/usr/local/go/bin/go mod init " + ProjectName
	// fmt.Println("change dir:", GenerateFolder, os.Chdir(GenerateFolder))

	// cmd := exec.Command(cli)
	// err := cmd.Run()
	// fmt.Println(err, GenerateFolder, cli)
}
