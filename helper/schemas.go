package helper

import (
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	log "github.com/sirupsen/logrus"
)

// prepare simplified list of schema, map by schema name
func PrepareSchemas() {
	for schemaname, schemasetting := range Doc.Components.Schemas {
		log.Info("Prepare Schema: ", schemaname)

		//Schema name "Error" require special field error_code, error_msg, need validation
		if schemaname == "Error" {
			ValidateIfErrorSchema(schemasetting)
		}
		requiredfields := schemasetting.Value.Required
		modelname := GetModelName(schemaname)
		//initiate new schema
		var schemaobj = Model_SchemaSetting{
			ModelName:     modelname,
			ModelType:     LowerCaseFirst(schemasetting.Value.Type),
			ModelFormat:   LowerCaseFirst(schemasetting.Value.Format),
			InterfaceName: GetInterfaceName(schemaname),
			Description:   schemasetting.Value.Description,
			FieldList:     map[string]Model_Field{},
		}
		//prepare properties of this schema
		for fieldname, fieldsettings := range schemasetting.Value.Properties {
			var fieldobj = Model_Field{}
			fieldobj.FieldIsModel = false
			fieldobj.ApiFieldName = fieldname
			fieldobj.FieldName = UpperCaseFirst(fieldname)
			fieldobj.Description = fieldsettings.Value.Description
			fieldobj.Example = getExamples(fieldsettings)

			// define this field belong to which model, easier for reference
			fieldobj.ModelName = modelname

			// openapi field format break into 2 properties,
			// we need both to define proper variable type, in diff programming language
			fieldobj.FieldType = fieldsettings.Value.Type
			if fieldobj.FieldType == "object" {
				fieldobj.FieldIsModel = true
				if fieldsettings.Ref != "" {
					log.Fatal("Schema ", schemaname, " has field ", fieldname, ":object using ref which is not supported yet: ", fieldsettings.Ref)
				} else if fieldsettings.Value.Properties == nil {
					log.Fatal("Schema ", schemaname, " has field ", fieldname, ":object but undefine properties")
				}
			} else if fieldobj.FieldType == "array" {
				if fieldsettings.Value.Items.Ref != "" {
					fieldobj.ChildItemType = GetModelNameFromRef(fieldsettings.Value.Items.Ref)
				} else if fieldsettings.Value.Items.Value.Type != "" {
					fieldobj.ChildItemType = fieldsettings.Value.Items.Value.Type
				} else {
					log.Fatal(fieldname + " using type:array, but undefine items.type or items.$ref")
				}

			}

			fieldobj.ApiFieldType = LowerCaseFirst(fieldsettings.Value.Type)
			fieldobj.FieldFormat = LowerCaseFirst(fieldsettings.Value.Format)

			for _, requiredfieldname := range requiredfields {
				if fieldobj.ApiFieldName == requiredfieldname {
					fieldobj.Required = true
				}
			}
			log.Debug("    ", fieldname, " : ", fieldobj.FieldType, "  //example: ", fieldobj.Example)
			// assign this field setting become 1 of the property in schema
			schemaobj.FieldList[fieldobj.FieldName] = fieldobj
		}
		if schemasetting.Value.AdditionalProperties != nil {
			log.Warn("*** ADDITIONAL PROPERTIES: ", schemaname, ": ", schemasetting.Value.AdditionalProperties.Value.Type)
		}

		//assign newschema into schema list
		AllSchemas[schemaname] = schemaobj
		log.Info("Complete schemas")

	}
	// 1 all schemas
	// schemaname = {modelname,type=object/array,fields>fieldsettings,descriptions}
}

func getExamples(op *openapi3.SchemaRef) string {
	val := ""
	vartype := ""
	fieldsettings := op.Value
	example := op.Value.Example
	if example != nil {

		if op.Value.Type == "array" {
			vartype := "[]" + op.Value.Items.Value.Type
			val = fmt.Sprintf("%#v", example)
			val = strings.Replace(val, "[]interface {}", vartype, -1)

			log.Warning("type:", vartype)
		} else {
			val = fmt.Sprintf("%#v", example)
		}

	} else {
		if fieldsettings.Type == "array" {
			/*

				if fieldobj.FieldType == "array" {
					if fieldsettings.Value.Items.Ref != "" {
						fieldobj.ChildItemType = GetTypeNameFromRef(fieldsettings.Value.Items.Ref)
					} else if fieldsettings.Value.Items.Value.Type != "" {
						fieldobj.ChildItemType = fieldsettings.Value.Items.Value.Type
					} else {
						log.Fatal(fieldname + " using type:array, but undefine items.type or items.$ref")
					}

				}

			*/
			// vartype := fmt.Sprintf("%T", example)
			if fieldsettings.Items.Ref != "" {
				vartype = GetModelNameFromRef(fieldsettings.Items.Ref)
			} else if fieldsettings.Items.Value.Type != "" {
				vartype = fieldsettings.Items.Value.Type
			}

			val = "[]" + vartype + "{}"

			// if op.Value.Items.Ref != "" && op.Value.Items.Value.Example != nil {
			// 	// vartype = GetTypeNameFromRef(op.Value.Items.Ref)
			// 	valtype := op.Value.Items.Value.Type
			// 	values := fmt.Sprintf("%#v", op.Value.Items.Value.Example)
			// 	val = "[-]" + valtype + "{" + values + "}"
			// } else {
			// 	val = "[]" + op.Value.Items.Value.Type
			// }
			// log.Debug("array type, sub item type: ", op.Value.Items.Ref, ", ", op.Value.Items.Value.Type)
		}

	}
	return val
}

func ValidateIfErrorSchema(schemasetting *openapi3.SchemaRef) {
	has_error_code := false
	has_error_msg := false

	for fieldname, _ := range schemasetting.Value.Properties {
		if fieldname == "err_code" {
			has_error_code = true
		}
		if fieldname == "err_msg" {
			has_error_msg = true
		}
	}

	if has_error_code == false {
		log.Fatal("Schema 'Error' is special schema require property err_code")
	}
	if has_error_msg == false {
		log.Fatal("Schema 'Error' is special schema require property err_msg")
	}
}
