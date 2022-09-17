# Introduction
The goal of this project is to allow us develop microservices in shortest time, design with low code concept. 

We know openapi is cool, we use swagger or postman for api documentation and design. However, it is redundant cause design api, develop api, and ensure api compliance will spend lot of time. 

In this project, instead of develop restful api, and match with openapi standard, we request developer design openapi, and generate full microservices from the .yaml. No code, no extra documentation, no headache.



# Project Status
This project still under P.O.C stage. Discussion welcome.

# How to use
1. clone this project to your home director
```bash
cd ~
mkdir golang
cd golang
git clone https://github.com/SIMITGROUP/openapigenerator.git
```
2. build this project
```bash
cd openapigenerator
go build .
```
3. prepare openapi v3 spec file into this folder (just let you know there is sample, you can skip this step)
```bash
cp sample/spec.yml .
```
4. Generate ***openapi/openapi.go***, ***openapi/schema.go***, ***openapi/userfunction.go***
```bash
./openapigenerator spec.yml 
```
edit ***openapi/userfunction.go*** to perform actual api execution


5. initiate new go microservices project
```bash

mkdir ~/golang/openapiserver
cd ~/golang/openapiserver
cp ../openapigenerator/spec.yml .
go mod init openapiserver
cp -a ../openapigenerator/openapi .
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
go get .
go build .
./openapiserver
```

8. Try your api according your spec
