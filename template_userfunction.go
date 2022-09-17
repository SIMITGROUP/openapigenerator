package main

var Data_userfunction = ""
var Temp_userfunction = `package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DummyFunction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "DummyFunction"})
}
##data##
`
