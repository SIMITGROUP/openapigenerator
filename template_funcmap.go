package main

var Data_funcMap = ""
var Temp_funcMap = `package openapi

import "github.com/gin-gonic/gin"

var funcMap = map[string]gin.HandlerFunc{
	"DummyFunction":DummyFunction,##data##
}`
