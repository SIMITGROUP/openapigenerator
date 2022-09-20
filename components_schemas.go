package main

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func prepareSchemas(schemas openapi3.Schemas) string {

	schemastr := ""
	for schemaname, setting := range schemas {
		// fmt.Println("schema:", schemaname, setting.Value.Title)
		props := setting.Value.Properties
		tmp := ""
		for field, fieldsetting := range props {
			tmp = tmp + getFieldSettingStr(field, *fieldsetting.Value) + "\n"
		}
		schemastr = schemastr + "\ntype " + GetModelName(schemaname) + " struct {\n" + tmp + "}\n"
	}
	Temp_schema := ""
	Data_schema := schemastr
	data := strings.Replace(Temp_schema, "##data##", Data_schema, -1)
	data = schemastr
	return data
}

func getFieldSettingStr(name string, s openapi3.Schema) string {
	fieldtype := s.Type
	prefix := "    "
	if s.Type == "integer" {
		if s.Format != "" {
			fieldtype = s.Format
		} else {
			fieldtype = "int32"
		}

	} else if s.Type == "object" {
		// fieldtype = " string //original is object"
		tmp := ""
		for subfieldname, subfieldsetting := range s.Properties {
			tmp = tmp + prefix + getFieldSettingStr(subfieldname, *subfieldsetting.Value) + "\n"
		}
		return prefix + name + " struct{\n" + tmp + prefix + "}"
	} else if s.Type == "array" {
		fieldtype = "[]" + s.Items.Value.Type
	}
	// fmt.Println(cases.Title(language.Und).String("goSAMples.dev is the best Go bLog in the world!"))
	// newname := cases.Title(language.Und).String(name)
	newname := upperCaseFirst(name)
	return prefix + newname + " " + fieldtype + " `json:\"" + name + "\"`"
}
