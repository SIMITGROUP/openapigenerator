package main

import (
	"os"
)

func writeFile(folder string, filename string, content string) {
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
