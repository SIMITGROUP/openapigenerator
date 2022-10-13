package buildgo

import "openapigenerator/helper"

func ConvertDataType(datatype string, dataformat string, fieldname string) string {
	dtype := ""

	if datatype == "integer" {
		if dataformat == "int64" {
			dtype = "int64"
		} else {
			dtype = "int32"
		}
	} else if datatype == "number" {
		if dataformat == "float" {
			dtype = "float32"
		} else {
			dtype = "float64"
		}
	} else if datatype == "boolean" {
		/*kiv seems string "byte","binary" not so relavent, ignore it*/
		dtype = "bool"
	} else if datatype == "string" {
		/*kiv  seems string "byte","binary","date","date-time","password" not so important,
		ignore format the the moment*/
		dtype = "string"
	} else if datatype == "array" {
		dtype = "array"
	} else if datatype == "object" {
		dtype = helper.GetModelName(fieldname)
	} else {

		/*limitation others typeall put as string at the moment*/
	}

	return dtype
}
