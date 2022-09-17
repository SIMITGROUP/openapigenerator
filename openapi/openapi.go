package openapi

import (
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

func Serve(docfile string, listen string) {
	router := gin.Default()
	doc, _ := openapi3.NewLoader().LoadFromFile(docfile)
	handleURL(doc, router)
	router.Run(listen)
}

/************ map url to operation function *************/
func handleURL(doc *openapi3.T, r *gin.Engine) gin.Engine {
	for oripath, pathmethods := range doc.Paths {
		path := convertGinPath(oripath)

		if pathmethods.Get != nil {
			r.GET(path, funcMap[pathmethods.Get.OperationID])
		}
		if pathmethods.Post != nil {
			r.POST(path, funcMap[pathmethods.Get.OperationID])
		}
		if pathmethods.Put != nil {
			r.PUT(path, funcMap[pathmethods.Get.OperationID])
		}
		if pathmethods.Delete != nil {
			r.DELETE(path, funcMap[pathmethods.Get.OperationID])
		}
	}
	return *r
}

/**************************** tools **************************/
func convertGinPath(oripath string) string {
	newpath := oripath
	r := regexp.MustCompile(`{\s*(.*?)\s*}`)
	matches := r.FindAllStringSubmatch(newpath, -1)
	for _, v := range matches {
		openapistr := "{" + v[1] + "}"
		reststr := ":" + v[1]
		newpath = strings.Replace(newpath, openapistr, reststr, -1)
	}

	return newpath
}

