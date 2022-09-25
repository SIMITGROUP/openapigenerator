package helper

import (
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Allroutes = []MethodSettings{}
var Allschemas = openapi3.Schemas{}
var Allhandles = []Model_Handle{}
var Allsecurityschemas = openapi3.SecuritySchemes{}
var Proj = ProjectSetting{}

func UpperCaseFirst(name string) string {
	newname := cases.Title(language.Und).String(name)
	return newname
}
func LowerCaseFirst(name string) string {
	newname := cases.Lower(language.Und).String(name)
	return newname
}

func GetAuthMethodName(schemename string) string {

	// for name, schemesetting := range Allsecurityschemas {
	// 	if name == schemename {
	// 		return "ss"
	// 	}
	// }
	return "Auth_" + schemename
}
func GetModelName(name string) string {
	return "Model_" + name
}

func GetInterfaceName(name string) string {
	return "Interface_" + name
}

func ReadFile(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func WriteFile(folder string, filename string, content string) {
	GenerateFolder := Proj.GenerateFolder
	targetfolder := ""
	targetfile := ""

	if folder != "" {
		targetfolder = GenerateFolder + "/" + folder
		targetfile = GenerateFolder + "/" + folder + "/" + filename
	} else {
		targetfolder = GenerateFolder
		targetfile = GenerateFolder + "/" + filename
	}

	_ = os.MkdirAll(targetfolder, 0777)
	// fmt.Println("targetfile:", GenerateFolder, targetfile, "===", targetfolder, err)
	_ = os.WriteFile(targetfile, []byte(content), 0644)
}

func ConvertGinPath(oripath string) string {
	newpath := oripath
	r := regexp.MustCompile(`{\s*(.*?)\s*}`)
	matches := r.FindAllStringSubmatch(newpath, -1)
	for _, v := range matches {
		openapistr := "{" + v[1] + "}"
		reststr := ":" + v[1]
		newpath = strings.Replace(newpath, openapistr, reststr, -1)
	}

	return newpath
}
func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GetModelNameFromRef(refstring string) string {
	typename := GetTypeNameFromRef(refstring)
	return GetModelName(typename)
}
func GetTypeNameFromRef(refstring string) string {
	refer_arr := strings.Split(refstring, "/")
	typename := refer_arr[len(refer_arr)-1]
	return typename
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
