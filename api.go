package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

func NewServer(docfile string) *gin.Engine {
	var authMiddleware *jwt.GinJWTMiddleware
	var err error
	r := gin.Default()

	//read api documents
	doc, _ := openapi3.NewLoader().LoadFromFile(docfile)

	prepareComponent(doc)
	preparePaths(doc)
	prepareHandles(doc)
	// handleURL(doc, router)
	// router.Run(listen)

	//prepare all components
	//schema
	//security

	//

	//prepare JWT token
	authMiddleware, err = OpenapiJWT()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	setDefaultRoute(r, authMiddleware)

	return r
}

func setDefaultRoute(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

}
