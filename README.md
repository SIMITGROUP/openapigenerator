# Openapi Generator
An openapiv3 generator which can generate micro-services for GO Language

*** this project good enough for prototyping, and the documentation is ongoing ***

## Content
- [Openapi Generator](#openapi-generator)
  - [Contents](#contents)
  - [Introduction](#introduction)
  - [Quickstart](#quickstart)
  - [Dockerizing](#dockerizing)
- [Openapi Document](#openapi-document)
    - [Requirements](#requirements)
    - [Define Environment Variables](#define-environment-variables)
    - [Define Error Schema](#define-error-schema)
    - [Require 2XX and 4XX In Every Response ](#require-2XX-and-4XX-in-every-response)    
    - [Not Supported Features](#not-supported-features)
- [Develop In Generated Project](#develop-in-generated-project)     
    - [Project Overview](#project-overview) 
        - [Development Concept](#development-concept)  
        - [File Structure](#file-structure) 
        - [Schemas Or Models](#Schemas-or-models) 
        - [Routes](#routes) 
        - [Route Handles](#route-handles) 
    - [Develop Real Route Handle](#develop-real-route-handle)
        - [Transfer Route Handle Into New File](transfer-route-handle-into-new-file)
        - [Define Handle Exists In Document](define-handle-exists-in-document)
        - [Create Environment Variables](#create-environment-variables)
    - [Build Project](#build-project) 
        - [Simple Build](#simple-build) 
        - [Build For Distribution](#build-for-distribution) 
        - [Build For Docker](#build-for-docker) 
    - [Unit Test](#unit-test)
        - [Define Real Test](#define-real-test) 
        - [Execute Unit Test](#execute-unit-test)   
    - [Secure Microservice](#secure-microservice) 
        - [Secure with ApiKey](#secure-with-apikey) 
- [Todo](#todo)


## Introduction
We know openapi is cool, usually we use swagger or postman for api testing, documentation and design. However, design, develop, testing and documentation is redundant work and spend us lot of time.

This generator able to generate micro-services with below capability:
1. working gin mock server, you can modified it to become real microservice
2. unit test for all restapi
3. build in swagger-ui to allow you test api easily
4. easily generate docker image
5. build in support security scheme


## Quickstart
Make sure you api document comply to [Requirements](#requirements), then you can generate micro-service source code:
1. [Download openapigenerator](https://github.com/SIMITGROUP/openapigenerator/releases)
2. prepare openapi-v3 spec file, (tested in .yaml only)
3. Execute below command :
```bash
./openapigenerator-mac.bin --apifile="simpleapi.yaml"  --targetfolder=myproject --projectname="project1" --port="9000"  --lang="go"
```
4. You can use Visual Studio Code or others IDE to open `~/myproject`:
```bash
cd myproject
code .
```

5. Copy .env.default to .env, define appropriate value, like API keys and etc
6. Run below command and done:
```bash
make && ./project1 
```
7. Try access rest api interface:
     http://localhost:9000/doc/swagger-ui/index.html



## Dockerizing
Project come with few command help you run microservice in docker environment. Check `Makefile` to know more detail

1. create docker container in devnetwork.
```
cd myproject
make dockerremove ; make docker &&  make dockerrun
make dockershell # access docker shell 
```
2 You may edit `Makefile` if you wish to pass create docker container with more useful setting


# Openapi Document
Pure openapi not good enough for prepare complete microservices, there is additional guide line shall follow at below section. 

## Requirements
This project has below requirement:
1. api document written in openapi v3
2. api document shall save in .yaml format
3. define operationId in every api in path, with value:
    * Alphabet value which can use to define as programming function title
    * No special characters allowed, except prefix with '-' (avoid generate route handle)
    * example: `GetData`, `SaveData`, `-SaveData`
4. All element shall refere from #components/schema, example:
    * parameters
    * request body
    * response content
## Not Supported Features
1. openconnect security schemes
2. apiKey only supported at header, others area is not supported (cookie, query...)

## Define Environment Variables
We can use [Specification Extensions](#https://spec.openapis.org/oas/v3.0.3#specificationExtensions) to declare environment variable for microservices, it will add into `.env.default`.

Example:
```yaml
x-env-vars:
  MONGO_SERVER: mongodbserver
  MONGO_DB: db1
  MONGO_USER: 
  MONGO_PASS: 
```
# Define Error Schema
You can define special schema "Error", which shall include 2 field: `err_code` and `err_msg` as below example.
 http request for `post`,`put` which you may submit requestbody.
```yaml
    Error:
      type: object      
      properties:
        err_code:
          type: string
          example: "ER-MG-001"
        err_msg:
          description: A human readable error message
          type: string
          example: "Mongodb not connected"
        version:
          type: string
          example: v1
```
How it work:
1. This schema can define as error 4xx response for http request with requestbody (post, put, patch)
2. when the actual request body not compatible with schema defination, the `err_msg` will display error like below:
```json
{
  "err_code": "ERR_INPUT_VALIDATION",
  "err_msg": "Key: 'Model_Book.Bookid' Error:Field validation for 'Bookid' failed on the 'required' tag",
  "version": "v1"
}
```

# Require 2XX and 4XX In Every Response
For good practise, every http request shall define at least 1 response for 2xx and 4xx. There is hardcoded behaviour which will identity 2xx as  success response, and 4xx is failed response.



# Develop In Generated Project
After code generated, we can start over development by change the generated code. You shall prepare below dev environment:
1. Visual Studio Code
2. Go Sdk
3. Docker (optional if you play with docker)

## Project Overview
The generated project prepare under below structure:
1. written in go language
2. Running on top of [gin](#https://github.com/gin-gonic/gin) web service
3. It work as complete mock server to service micro-service as stated in openapi document.

### Development Concept
To make the micro-service perform real task, we shall perform development, which involve below scope:
1. Modify route handle to perform real task:
    * `route` is restapi request like `GET /myapi`, `POST /myapi/res1`
    * `route handle` is programming function, which will trigger when client access specific `route`. 
    * Every `route` connect to own `route handle`
    * `route` is terminology from `gin`, it have similar meaning with openapi `path`

2. Unit test (optional)
3. Prepare environment for deployment like:
    * environment variables in `.env` or `.env.docker`
    * prepre dockers images
    * prepare database for handle

On and off, depends on project requirement you need to perform modification at openapi document, and regenerate the code again, and again. It is important to know how to regenerate the code without overwrite your modification. 

Technically it has few rules:
1. Remain all file with prefix Z*.go remain unchange (openapi/Z*.go test/Z*.go)
2. During development, create new `.go` file in folder `openapi/`
3. Created file should not start with `Z`, to avoid it mixed with generated file 
4. Root level file as below guideline:
    * `go.main`: always overwrite by generator
    * `go.mod`:  always overwrite by generator
    * `Makefile`:  always overwrite by generator
    * `.env.default`: always overwrite by generator
    * `.env`: free to change    
    * `.env.docker`: free to change
    
### File Structure
Below is some explanation generated code
* `main.go`: entry point
* `go.mod`: go project setting
* `.env.default`: template for `.env`
* `Makefike`: store command for build project, you may change if you have complex requirement
* `Dockerfile`: template for generate docker image, change it if you have special need on docker
* `openapi/`: store all go program, you shall perform develop in this folder only
* `openapi/ZModel_*.go`: All openapi schema will generate as go `Model` here (structs)
* `openapi/ZRouterRegistry.go`: List all the route/path in this microservice, and connect to which route handle
* `openapi/ZRouterHandle.go`: Keep all default auto generate `route handle`, base on name defined in openapi's path's operationId. 
* `openapi/ZSecurity*.go`: security schemes setting according openapi document
* `openapi/ZServer.go`: store code to run gin server
* `api/api.yaml`: store openapi document, used by swagger-ui
* `dist/*`: project binary build into this folder

### Schemas Or Models
Openapi's schema equivalent to `model` in this project, and all model store as `openapi/ZModel_*.go`. In go, `Model` is kind of struct, suitable interface, getter/setter/validator was prepared.

The model is important cause it act as pattern of api output. Route handle comply output pattern using specific `model`.

### Routes
`Route` equivalent to rest api `request method` + `path`, example
* Get /api/service1
* Post /api/v1/service2
* Put /api
* Delete /api/crud

### Route Handles
1. `Route handle` is programming `function` like `GetStudents()` which is trigger by `route`
2. openapi document `operationId` declare name of `route handle`, and generator will help you prepare dummy route handle. 

## Develop Real Route Handle
To provide real data in micro-service, we perform 3 step below:
1. [Transfer Route Handle Into New File](transfer-route-handle-into-new-file)
2. [Define Handle Exists In Document](define-handle-exists-in-document)
3. [Create Environment Variables](#create-environment-variables)

### Transfer Route Handle Into New File
1. Create new file `openapi/routehandles.go` (or others name you like)
2. `Cut and paste` specific function from `openapi/ZRouteHandle.go` into `openapi/routehandles.go`. Let's assume `getMemoryInfo(c *gin.Context){}`
    * remain `getMemoryInfo()` in `ZRouteHandle.go` will cause duplicate function and cause error
3. Edit content of `openapi/routehandles.go` to serve real data:
```go
// auto generate by generator
package openapi

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	// "net/http"
)


func getMemoryInfo(c *gin.Context) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	allocated := bToMb(m.TotalAlloc)
	available := bToMb(m.Sys)
	percent := int((allocated/available)*100) / 100
	allocatedstr := fmt.Sprintf("%v MB", allocated)
	availablestr := fmt.Sprintf("%v MB", available)
	percentstr := fmt.Sprintf("%v percent", percent)
	c.Header("Content-Type", "application/json")
	// 
	// type Model_MemoriesInfo struct{
	// 	Percent string `json:"percent" binding:""` //
	// 	Total string `json:"total" binding:""` //
	// 	Used string `json:"used" binding:""` //
    //  Version string `json:"version"`
	// }

	data := Model_MemoriesInfo{} //Model_MemoriesInfo defined at openapi/ZModel_MemoriesInfo.go
	data.SetTotal(percentstr)
	data.SetPercent(availablestr)
	data.SetUsed(allocatedstr)
    data.SetVersion("v1")
	c.JSON(200, data)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

```
### Define Handle Exists In Document
1. Edit the original openapi document, add entry `x-operationId-exists`:
```yaml
# bottom of file
x-operationId-exists:
  getMemoryInfo: true
```
2. Try regenerate the source code, `getMemoryInfo()` no longer generate in `ZRouteHandle.go`

### Create Environment Variables
Lot of time, we wish to use configure microservice according deployment requirement, such as:
1. Define database server location and credentials
2. On/Off specific functions
3. Define apikey and etc


1. Define environment variable in openapi document using `x-env-vars` and regenerate the source code. 
Example in openapi document:
```yaml
# bottom of file
x-operationId-exists:
  getMemoryInfo: true
x-env-vars:
  APIVERSION: v1.1
```
Generated `.env.default` will automatically add below entry
```env
APIVERSION=v1.1
```
2. load environment variable in route handle using `godotenv.load()` and `os.Getenv("APIVERSION")`. Example of updated code:
```go

package openapi

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	// "net/http"
)

func getMemoryInfo(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	verionno := os.Getenv("APIVERSION")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	allocated := bToMb(m.TotalAlloc)
	available := bToMb(m.Sys)
	percent := int((allocated/available)*100) / 100
	allocatedstr := fmt.Sprintf("%v MB", allocated)
	availablestr := fmt.Sprintf("%v MB", available)
	percentstr := fmt.Sprintf("%v percent", percent)
	c.Header("Content-Type", "application/json")
	//
	// type Model_MemoriesInfo struct{
	// 	Percent string `json:"percent" binding:""` //
	// 	Total string `json:"total" binding:""` //
	// 	Used string `json:"used" binding:""` //
	// }

	data := Model_MemoriesInfo{} //Model_MemoriesInfo defined at openapi/ZModel_MemoriesInfo.go
	data.SetTotal(percentstr)
	data.SetPercent(availablestr)
	data.SetUsed(allocatedstr)
	data.SetVersion(verionno)
	c.JSON(200, data)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

```

3. define environment variable via:
    * edit `.env` in non-docker build
    * edit `.env.docker` in docker build
    * or, modify it in command line, `export X_API_Key=12345` 

## Build Project
We shall build the project to run the micro-service, follow below instruction.

### Simple Build
1. run below command:
```bash
make
```
2. To run the service:
Linux/Mac:
```bash
./<your-project-name>
```
Windows: double click the filename `<your-project-name>`
### Build For Distribution
If you wish to build binary for specific OS and distribute manually. Folow step
1. Run below command:
```bash
make windows #build for windows
make linux #build for linux
make mac #build for mac
make mac-arm #build for mac m1 type processor

```
2. Distrubute file in `dist/`, along with suitable `.env` file

### Build For Docker
We can build and run docker image via below step:
1. Copy `.env.default` to `.env.docker`, and define appropriate content
2. run command:
```bash
make dockerremove #remove existing docker 
make docker # create docker image only
make dockerrun # create docker container
```
or we can simplified as single row command:
```bash
make dockerremove; make docker && make dockerrun
```
3. We can change Docker image build from different source by edit `Dockerfile`. Like replace `alpine` to `ubuntu`
4. We can change docker image/container/network by edit `Makefile`

## Unit Test
Generator help you prepare reasonable structure for unit test, however you need to copy content into suitable file to prevent effort overwritten when regenerate code.
You need to install `grc` package for unit test
https://jakeholmquist.medium.com/add-some-fun-to-your-cli-with-grc-ea868df985b6


### Define Real Test
1. expand folder ./test, open any file (assume ZGet_Memory_test.go)
2. scroll down and follow comment, copy content into test/Get_Memory.go (File name shall follow comment)
```go
// copy and modify below content and put into new file Get_Memory.go (in this test folder)
/*
package test
import (
	"io"
...
```
3. after step 2, error in ZGet_Memory_test.go should disappear
4. Modify `Get_Memory.go` if neccessary to serve real content
5. repeat same step 1-4 for others file until all error disappear:
    * if request body is required, `FunctionName_RequestBody()` will fill in sample data which defined in openapi document
    * You can change whole structure of unit test as long as the original function name remain
    * Don't change any file start with Z*.go, cause it will overwritten by generator

### Execute Unit Test
Ensure microservice is activated, run below command and see the result:
```bash
make apitest
```

## Secure Microservice
We can secure restapi via define secruityScheme in openapi document. 

### Secure With Apikey
1. Add security Schemes. Below is apiKey example:
```yaml
components:
  securitySchemes:
    BasicApiKey:
      type: apiKey
      in: header
      name: X-Api-Key #environment variable X_Api_Key will prepare automatically
```
2. Define which api use it:
```yaml
paths:
  /api1:
    get:
      summary: welcome
      description: show msg undefine resource
      operationId: "welcome"
      security:
        - BasicApiKey: []  # Get /api1 require X-Api-Key defined in header
```

3. Re-generate the source code:
    * `X_Api_Key` will automatically prepare in `.env.default`
    * `openapi/ZSecurity_BasicApiKeyy.go` generated
    * unit test template use env var X_Api_Key in header
4. Update `.env` and `.env.docker` if both file exists, and restart the micro-service
5. Try the `Get /api` again with header `X-Api-Key` as:
```bash
http get localhost:<portno>/api1 X-Api-Key:<you-key-code>
```

# Todo
1. Add `x-generator-setting` to direct set project name, port number and etc suitable configuration
2. rename this project to prevent crash name with official openapi-generator
3. support more component type
4. client generators for different kind of languages
5. add some common template for
   * crud for different kind of database
   * messaging template for sms, email, push notifications
6. openapi 3.1
7. more securityschemes
8. data validation


1. fix no .env.docker