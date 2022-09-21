package main

import (
	"fmt"

	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

func prepareSchemas(schemas openapi3.Schemas) string {

	for schemaname, setting := range schemas {
		// fmt.Println("schema:", schemaname, setting.Value.Title)
		modelname := GetModelName(schemaname)
		interfacename := GetInterfaceName(schemaname)
		props := setting.Value.Properties
		tmp := ""
		gettersetterstr := ""
		interfacecontent := "\n    Validate()"
		fmt.Println("prepare model", modelname, setting.Value.Type)
		//no properties, visit reference insteads
		if setting.Value.Type == "object" {
			for field, fieldsetting := range props {
				tmp = tmp + getFieldSettingStr(field, *fieldsetting.Value) + " // " + fieldsetting.Value.Description + "\n"

				golangfieldtype := convGoLangType(*fieldsetting.Value)
				intstr, getsetstr := retrieveGetSetStr(modelname, field, golangfieldtype)
				interfacecontent = interfacecontent + intstr
				gettersetterstr = gettersetterstr + getsetstr
			}
		} else if setting.Value.Type == "array" {
			//try do nothing for array
			continue
		} else {
			//no need create models/interface
			// return ""
			if props == nil {

				tmp = "    " + modelname + " []" + getModelNameFromRef(setting.Value.Items.Ref) + " `json:\"" + getTypeNameFromRef(setting.Value.Items.Ref) + "\"`\n"
			}
		}

		interfacestr := fmt.Sprintf("type %v interface {%v\n}", interfacename, interfacecontent)

		schemastr := "\ntype " + GetModelName(schemaname) + " struct {\n" + tmp + "}\n" +
			interfacestr + "\n\n" + getValidateStr(modelname) + "\n\n" + gettersetterstr
		content := `package openapi
import (
	// "regexp"
	// "github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)` + schemastr
		writeFile("openapi", modelname+".go", content)
	}

	return "" //schemastr
}
func getValidateStr(name string) string {
	// Validate implements basic validation for this model
	validationtemplate := `
func (m %v) Validate() error {
	return validation.Errors{
		//"name": validation.Validate(m.Name, validation.Required, validation.Length(5, 20)),
		//"email": validation.Validate(m.Name, validation.Required, is.Email),
		//"zip": validation.Validate(m.Address.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	}.Filter()
}`

	validationstr := fmt.Sprintf(validationtemplate, name)

	return validationstr
}

func retrieveGetSetStr(modelname string, orifieldname string, fieldtype string) (string, string) {
	fieldname := upperCaseFirst(orifieldname)
	gettemplate := `func (m %v) Get%v() %v {
	return m.%v
}`
	settemplate := `func (m *%v) Set%v(val %v) {
	m.%v = val
}`
	getstr := fmt.Sprintf(gettemplate, modelname, fieldname, fieldtype, fieldname)
	setstr := fmt.Sprintf(settemplate, modelname, fieldname, fieldtype, fieldname)
	getsetstr := getstr + "\n\n" + setstr + "\n\n"
	interfacestr := "\n    Get" + fieldname + "()\n    Set" + fieldname + "(" + fieldtype + ")"
	return interfacestr, getsetstr
}

func getFieldSettingStr(name string, s openapi3.Schema) string {
	newname := upperCaseFirst(name)
	fieldtype := convGoLangType(s)
	// ("sdds")

	fmt.Println("  after converttype =", newname, "=", fieldtype)
	prefix := "    "
	// if s.Type == "integer" {
	// 	if s.Format != "" {
	// 		fieldtype = s.Format
	// 	} else {
	// 		fieldtype = "int32"
	// 	}

	// } else if s.Type == "object" {
	// 	// fieldtype = " string //original is object"
	// 	tmp := ""
	// 	for subfieldname, subfieldsetting := range s.Properties {
	// 		tmp = tmp + prefix + getFieldSettingStr(subfieldname, *subfieldsetting.Value) + "\n"
	// 	}
	// 	return prefix + newname + " struct{\n" + tmp + prefix + "}"
	// } else if s.Type == "array" {
	// 	fieldtype = "[]" + s.Items.Value.Type
	// } else if s.Type == "string" {
	// 	//do nothing for string
	// } else if s.Type == "" { //custom types, refer to another type
	// 	refer_arr := strings.Split(s.Items.Ref, "/")
	// 	refermodel := GetModelName(refer_arr[len(refer_arr)-1]) // get Model name
	// 	// fmt.Println("check data:", name, refermodel)
	// 	fieldtype = refermodel
	// }
	// fmt.Println(cases.Title(language.Und).String("goSAMples.dev is the best Go bLog in the world!"))
	// newname := cases.Title(language.Und).String(name)

	return prefix + newname + " " + fieldtype + " `json:\"" + name + "\"`"
}

func convertToGoFieldType(fieldtype string, fieldformat string) string {
	if fieldtype == "integer" {
		if fieldformat != "" {
			fieldtype = fieldformat
		} else {
			fieldtype = "int32"
		}

	}
	return fieldtype
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
		if s.Items.Value.Type != "" {
			fieldtype = s.Items.Value.Type
		} else { // use custom type
			fieldtype = getModelNameFromRef(s.Items.Ref)
		}
		fieldtype = "[]" + fieldtype
	} else if s.Type == "" { //custom types, refer to another type
		fieldtype = getModelNameFromRef(s.Items.Ref)
		// refer_arr := strings.Split(s.Items.Ref, "/")
		// refermodel := GetModelName(refer_arr[len(refer_arr)-1]) // get Model name
		// // fmt.Println("check data:", name, refermodel)
		// fieldtype = refermodel
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

func getModelNameFromRef(refstring string) string {
	typename := getTypeNameFromRef(refstring)
	return GetModelName(typename)
}
func getTypeNameFromRef(refstring string) string {
	refer_arr := strings.Split(refstring, "/")
	typename := refer_arr[len(refer_arr)-1]
	return typename
}

// func getCustomTypeName(refstring string) string {

// 	refermodelname := GetModelName(typename) // get Model name
// 	return refermodelname
// }
