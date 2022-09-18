package main

var Data_interface = ""
var Temp_interface = `package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,##data##
}`
