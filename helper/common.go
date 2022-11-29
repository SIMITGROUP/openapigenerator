package helper

import (
	"embed"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var EmbedFiles embed.FS

func SetEmbedFiles(file embed.FS) {
	EmbedFiles = file
}
func UpperCaseFirst(name string) string {
	newname := cases.Title(language.Und).String(name)
	return newname
}
func LowerCaseFirst(name string) string {
	newname := cases.Lower(language.Und).String(name)
	return newname
}
func ReadFile(filename string) string {

	// data, err := Box.FindString("|")
	// Box.ResolutionDir = filename

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, _ := os.ReadFile(filename)
	// entries, err := EmbedFiles.ReadDir("templates/go")
	// fmt.Println(entries, err)
	data, _ := EmbedFiles.ReadFile(filename)

	// fmt.Println("filename: ", filename)
	// fmt.Println("    content", string(data))
	return string(data)
}

func CheckFileExists(folder string, filename string) bool {
	targetfile := Proj.GenerateFolder
	if folder == "" {
		targetfile = targetfile + "/" + filename
	} else {
		targetfile = targetfile + "/" + folder + "/" + filename
	}

	_, err := os.Stat(targetfile)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}

}
func WriteFile(folder string, filename string, content string) {
	GenerateFolder := Proj.GenerateFolder
	// targetfolder := ""
	targetfile := ""

	if folder != "" {
		// targetfolder = GenerateFolder + "/" + folder
		targetfile = GenerateFolder + "/" + folder + "/" + filename
	} else {
		// targetfolder = GenerateFolder
		targetfile = GenerateFolder + "/" + filename
	}
	PrepareFolder(targetfile)
	// _ = os.MkdirAll(targetfolder, 0777)
	// fmt.Println("targetfile:", GenerateFolder, targetfile, "===", targetfolder, err)
	err := os.WriteFile(targetfile, []byte(content), 0644)
	if err == nil {
		log.Info("Write file ", targetfile)

	} else {
		errormsg := fmt.Sprintf("Can't write file %v error %v", targetfile, err)
		log.Fatal(errormsg)

	}

}
func PrepareFolder(targetfile string) error {
	dir := filepath.Dir(targetfile)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatal("Can't write :", err.Error())
	}
	return err
}

func CopyFile(sourceFile string, destinationFile string) error {

	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	PrepareFolder(destinationFile)
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		log.Error("Error creating", destinationFile)
		log.Fatal(err)
		return err
	}
	return nil
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
	return schemename + "_Auth"
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

func ConvertPathParasCurlyToColon(oripath string) string {
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

func VerifyKeyname(s string) bool {
	for _, r := range s {
		if r == '_' || r == '-' {
			return true
		} else if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func In_array(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Left(txt string, no int) string {
	return SubString(txt, 0, no)

}
func Right(txt string, no int) string {
	return SubString(txt, no, len(txt))

}
func SubString(txt string, startno int, endno int) string {
	runes := []rune(txt)
	safeSubstring := string(runes[startno:endno])

	return safeSubstring
}

func ConvertApiKeyToEnvVar(keyname string) string {
	return strings.Replace(keyname, "-", "_", -1)
}
