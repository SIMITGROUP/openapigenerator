package main

import (
	"flag"
	"fmt"
)

var GenerateFolder = *flag.String("targetfolder", "../openapiserverfolder", "")
var ProjectName = *flag.String("projectname", "openapiserver", "")
var Apifile = *flag.String("apifile", "spec.yaml", "")
var Defaultport = *flag.Int("port", 8000, "listen port")

func main() {

	fmt.Println("GenerateFolder", GenerateFolder)
	fmt.Println(ProjectName, GenerateFolder, Apifile)
	Generate(Apifile)

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
