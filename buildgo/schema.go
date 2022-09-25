package buildgo

import (
	"bytes"
	"fmt"
	"openapigenerator/helper"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"

	"github.com/getkin/kin-openapi/openapi3"
)

func WriteSchemas() {
	// Model_Field Model_Schema

	for schemaname, setting := range helper.Allschemas {
		schemaobj := helper.Model_Schema{}
		modelname := helper.GetModelName(schemaname)
		interfacename := helper.GetInterfaceName(schemaname)

		schemaobj.ModelName = modelname
		schemaobj.InterfaceName = interfacename

		props := setting.Value.Properties
		log.Info("Prepare Schema: ", modelname, "(", schemaname, ")", ": ", setting.Value.Type)
		if setting.Value.Type == "object" {
			allfields := []helper.Model_Field{}
			for field, fieldsetting := range props {
				// fmt.Println("Schema:", schemaname, field)
				examplestr := ""
				if fieldsetting.Value.Example == nil {
					if fieldsetting.Value.Type == "object" && fieldsetting.Value.Items.Ref == "" {
						log.Fatal("Schema " + schemaname + "." + field + " type=object, but not ref to another schema")
					}

					if fieldsetting.Value.Type != "" && fieldsetting.Value.Type != "object" && fieldsetting.Value.Type != "array" {
						log.Fatal("Undefine sample data in schema '" + schemaname + "' field '" + field + "'")
					}
					// fmt.Println("field ", field, fieldsetting.Value.Items)
					if fieldsetting.Value.Items.Ref != "" {
						fmt.Println(field, " == ", fieldsetting.Value.Items.Ref)
						examplestr = helper.GetModelNameFromRef(fieldsetting.Value.Items.Ref) + "{}"
						if fieldsetting.Value.Type == "array" {
							examplestr = "[]" + examplestr
						}

					}

				} else {
					examplestr = fmt.Sprintf("%#v", fieldsetting.Value.Example)
					examplestr = strings.Replace(examplestr, "interface {}", "string", -1)
				}
				fieldname := helper.UpperCaseFirst(field)
				fieldtype := convGoLangType(*fieldsetting.Value)
				log.Debug("    ", fieldname, ", ", fieldtype)
				fieldobj := helper.Model_Field{
					ModelName:    modelname,
					FieldName:    fieldname,
					FieldType:    fieldtype,
					ApiFieldName: field,
					Description:  fieldsetting.Value.Description,
					Example:      examplestr,
				}
				allfields = append(allfields, fieldobj)
				// _, _ = field, fieldtype
			}
			schemaobj.FieldList = allfields
		} else if setting.Value.Type == "array" { //array no need new model
			continue
		} else {

		}
		_, _, _, _ = schemaname, modelname, interfacename, props

		var writebuffer bytes.Buffer
		filename := modelname + ".go"
		schemapath := "./templates/go/schema.gotxt"
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
