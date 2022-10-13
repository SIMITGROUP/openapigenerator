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

		modelname := GetModelName(schemaname)
		//initiate new schema
		var schemaobj = Model_SchemaSetting{
			ModelName:     modelname,
			ModelType:     LowerCaseFirst(schemasetting.Value.Type),
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
				if fieldsettings.Value.Properties == nil {
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
			log.Debug("    ", fieldname, " : ", fieldobj.FieldType, "  //example: ", fieldobj.Example)
			// assign this field setting become 1 of the property in schema
			schemaobj.FieldList[fieldobj.FieldName] = fieldobj
		}
		//assign newschema into schema list
		AllSchemas[schemaname] = schemaobj

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
