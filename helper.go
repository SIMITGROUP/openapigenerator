package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func upperCaseFirst(name string) string {
	newname := cases.Title(language.Und).String(name)
	return newname
}
