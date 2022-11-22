package buildgo

import (
	"bytes"
	"openapigenerator/helper"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/getkin/kin-openapi/openapi3"
)

func WriteSchemas() {
	// Model_Field Model_Schema

	for schemaname, schemaobj := range helper.AllSchemas {
		log.Debug("Write Model ", schemaobj.ModelName, ": ", schemaobj.ModelType, " (", schemaname, ")")

		for f, fsetting := range schemaobj.FieldList {

			ftype := fsetting.FieldType
			format := fsetting.FieldFormat
			newtype := ConvertDataType(ftype, format, f)
			log.Debug("    Prepare field ", f, ": ", ftype, ":", newtype, " //", fsetting.ChildItemType)
			fsetting.FieldType = newtype
			schemaobj.FieldList[f] = fsetting

		}

		var writebuffer bytes.Buffer
		filename := "Z" + schemaobj.ModelName + ".go"
		schemapath := "templates/go/schema.gotxt"
		schemasrc := helper.ReadFile(schemapath)
		schematemplate := template.New("schema")
		schematemplate, _ = schematemplate.Parse(schemasrc)

		err := schematemplate.Execute(&writebuffer, schemaobj)
		if err != nil {
			log.Fatal("writing template ", filename, "error, ", err)
		}
		helper.WriteFile("openapi", filename, writebuffer.String())
	}
}

func convGoLangType(s openapi3.Schema) string {
	fieldtype := s.Type

	if s.Type == "integer" {
		if s.Format != "" {
			fieldtype = s.Format
		} else {
			fieldtype = "int32"
		}

	} else if s.Type == "string" {
		//do nothing for string
	} else if s.Type == "array" {
		if s.Items.Value.Type != "" && s.Items.Value.Type != "object" {
			fieldtype = s.Items.Value.Type
		} else { // use custom type
			fieldtype = helper.GetModelNameFromRef(s.Items.Ref)
		}
		fieldtype = "[]" + fieldtype
		log.Debug("It is array")
	} else if s.Type == "" { //custom types, refer to another type
		fieldtype = helper.GetModelNameFromRef(s.Items.Ref)
		// refer_arr := strings.Split(s.Items.Ref, "/")
		// refermodel := GetModelName(refer_arr[len(refer_arr)-1]) // get Model name
		// // fmt.Println("check data:", name, refermodel)
		// fieldtype = refermodel
	} else {
		//all others not supported treat as string
		fieldtype = "string"
	}
	/* else if s.Type == "object" {
		// fieldtype = " string //original is object"
		tmp := ""
		for subfieldname, subfieldsetting := range s.Properties {
			tmp = tmp + prefix + getFieldSettingStr(subfieldname, *subfieldsetting.Value) + "\n"
		}
		return prefix + newname + " struct{\n" + tmp + prefix + "}"
	}
	*/
	return fieldtype
}
