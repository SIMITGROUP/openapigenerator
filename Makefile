run:
	go get .
	go build .
	@echo "Build successfully, run command with your own options"
	@echo './openapigenerator  --targetfolder="../openapiserverfolder" --projectname="openapiserver" --listen=":9000" --apifile="samples/spec.yaml" --lang="go"'
windows:
	go get .
	GOOS=windows GOARCH=amd64 go build -o dist/openapigenerator-win.exe .
linux:
	go get .
	GOOS=linux GOARCH=amd64 go build -o dist/openapigenerator-linux.bin .
mac:
	go get .
	GOOS=darwin GOARCH=amd64 go build -o dist/openapigenerator-mac.bin .