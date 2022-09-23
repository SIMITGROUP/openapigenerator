run:
	go get .
	go build .
	@echo "Build successfully, run command with your own options"
	@echo './openapigenerator  --targetfolder="../openapiserverfolder" --projectname="openapiserver" --listen=":9000" --apifile="samples/spec.yaml" --lang="go"'