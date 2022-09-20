package main

import (
	"fmt"
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

	err := os.MkdirAll(targetfolder, 0777)
	fmt.Println("targetfile:", GenerateFolder, targetfile, "===", targetfolder, err)
	_ = os.WriteFile(targetfile, []byte(content), 0644)
}
