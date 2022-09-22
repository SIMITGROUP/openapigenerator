package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func upperCaseFirst(name string) string {
	newname := cases.Title(language.Und).String(name)
	return newname
}
func lowerCaseFirst(name string) string {
	newname := cases.Lower(language.Und).String(name)
	return newname
}

func GetAuthMethodName(name string) string {
	return "Auth_" + name
}
func GetModelName(name string) string {
	return "Model_" + name
}

func GetInterfaceName(name string) string {
	return "Interface_" + name
}
