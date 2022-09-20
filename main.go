package main

import (
	"flag"
)

var GenerateFolder = "" //*flag.String("targetfolder", "../openapiserverfolder", "")
var ProjectName = ""    //*flag.String("projectname", "openapiserver", "")
var ApiFile = ""        // *flag.String("apifile", "spec.yaml", "")
var ListenAddress = ""  // *flag.String("listen", ":8989", "listen address")

func init() {

	// GenerateFolder = *flag.String("targetfolder", "../openapiserverfolder", "")
	// ProjectName = *flag.String("projectname", "openapiserver", "")
	// ApiFile = *flag.String("apifile", "samples/spec.yaml", "")
	// ListenAddress = *flag.String("listen", ":8982", "listen address")

}
func main() {

	flag.StringVar(&GenerateFolder, "targetfolder", "../openapiserverfolder", "Generate Folder to which folder")
	flag.StringVar(&ProjectName, "projectname", "openapiserver", "Rest API GO project name")
	flag.StringVar(&ApiFile, "apifile", "spec.yaml", "openapi v3 yaml file")
	flag.StringVar(&ListenAddress, "listen", ":8982", "listen to which address, default :8982")

	flag.Parse()

	Generate(ApiFile)

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	// errInit := authMiddleware.MiddlewareInit()

	// if errInit != nil {
	// 	log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	// }

	// r.POST("/login", authMiddleware.LoginHandler)

	// auth := r.Group("/auth")
	// // Refresh time can be longer than token timeout
	// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// auth.Use(authMiddleware.MiddlewareFunc())
	// {
	// 	auth.GET("/hello", helloHandler)
	// }

	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }
	// r.Run(port)
}
