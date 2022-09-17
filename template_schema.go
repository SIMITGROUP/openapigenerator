package main

var Data_schema = ""
var Temp_schema = `package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,##data##
}`
