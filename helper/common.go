package helper

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func UpperCaseFirst(name string) string {
	newname := cases.Title(language.Und).String(name)
	return newname
}
func LowerCaseFirst(name string) string {
	newname := cases.Lower(language.Und).String(name)
	return newname
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
	err := os.WriteFile(targetfile, []byte(content), 0644)
	if err == nil {
		log.Info("Write file ", targetfile)

	} else {
		errormsg := fmt.Sprintf("Can't write file %v error %v", targetfile, err)
		log.Fatal(errormsg)

	}

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

func GetModelNameFromRef(refstring string) string {
	typename := GetTypeNameFromRef(refstring)
	return GetModelName(typename)
}
func GetSchemaFromRef(refstr string) Model_SchemaSetting {
	refer_arr := strings.Split(refstr, "/")
	schemaname := refer_arr[len(refer_arr)-1]
	return AllSchemas[schemaname]
}
func GetTypeNameFromRef(refstring string) string {
	refer_arr := strings.Split(refstring, "/")
	typename := refer_arr[len(refer_arr)-1]
	return typename
}
