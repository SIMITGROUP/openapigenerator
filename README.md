# Introduction
This project generate openapi3 server code for gin restapi

# How to use
1. clone this project to your home director
```bash
cd ~
git clone https://github.com/SIMITGROUP/openapigenerator.git
```
2. build this project
```bash
cd ~/openapigenerator
go build .
```
3. copy openapi v3 spec file into this folder
```bash
cp ~/spec.yml ~/openapigenerator
```
4. Generate openapi/openapi.go, openapi/schema.go, openapi/userfunction.go
```bash
./openapigenerator spec.yml 
```
5. initiate new go microservices project
```bash

mkdir ~/openapiserver
cd ~/openapiserver
cp ~/openapigenerator/spec.yml ~/openapiserver
go mod init openapiserver
cp -a ../openapigenerator/openapi ~/openapiserver
touch main.go
```

6. Edit main.go
```go
package main

import (
	"openapiserver/openapi"
)

func main() {

	openapi.Serve("spec.yml", ":8080")

}

```

7. run server
```bash
cd ~/openapiserver
go run .
```

8. Try your api according your spec
