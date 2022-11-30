# Openapi Generator
An openapiv3 generator which can generate micro-services for GO Language

## Content
<hr/>

- [Gin Web Framework](#gin-web-framework)
  - [Contents](#contents)
  - [Introduction](#introduction)
  - [Quick Start](#quickstart)



## Introduction
We know openapi is cool, usually we use swagger or postman for api testing, documentation and design. However, design, develop, testing and documentation is redundant work and spend us lot of time.

This generator able to generate micro-services with below capability:
1. working gin mock server, you can modified it to become real microservice
2. unit test for all restapi
3. build in swagger-ui to allow you test api easily
4. easily generate docker image
5. build in support security scheme


## Quickstart
1. Download openapigenerator
2. prepare openapi-v3 spec file, (tested in .yaml only)
3. Execute below command:
```bash
./openapigenerator --apifile="openapiv3.yaml"  --targetfolder="~/myproject" --projectname="project1" --port="9000"  --lang="go"
```
4. Use vscode to open `~/myproject`




## Setup
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
make
./openapigenerator --apifile="samples/spec.yaml"  --targetfolder="../myproject" --projectname="project1" --port="9000"  --lang="go"
```

3. use your rest api
    3.1 copy myproject/.env.default to .env
    3.2 fill in suitable info into .env
    3.3 run below command
```bash
cd ../myproject
make
./project1
```

4. Try your rest api http://localhost:9000, to access your mock rest api server. It will run return sample data defined in .yaml file.

5. Browse to http://localhost:9000/doc/swagger-ui/index.html to access swagger-ui

6. You can perform automatic unit test by open another terminal, run following command:
```bash
make apitest
```

## Run in Docker
Project come with few command help you run microservice in docker environment. Check `Makefile` to know more detail

1. create docker container in devnetwork.
```
cd myproject
make dockerremove ; make docker;  make dockerrun #run docker, under devnetwork
make dockershell # access docker shell 
```
2 You may edit `Makefile` if you wish to pass create docker container with more useful setting



# Features
1. Auto prepare model/interface according each component's schema
2. Auto prepare path and route to handles according operationID (required), response example data according reference schema and examples
3. Supported http traffic (GET/POST/PUT/DELETE)
4. Build in gin http server
5. Work with application/json response (only)
6. Use middleware control security requirement (support basic and apikey at this moment)
7. Supported Component type
    schema
    securityscheme


# Todo
1. add in missing basic and bearer jwt authorization
2. prepare log system
3. auto generate unit test
4. missing support component type
    Parameters
    RequestBodies
        connect still connect to schema, 
            require: true/false
            application/json only
    Responses
    Examples
    Callbacks
    Headers
    Links
5. Schema extends another Schema
6. Array of Object Schema
7. Different example but same schema (Example is put at response / requestBody)
8. Header name Xxx-Xxx-Xxx variable convert to xxx_Xxx_Xxx in code.
9. enum

# Rules while using this project
1. Not support  ***oneOf, anyOf, allOf, not ***
2. security scheme for apikey, oauth2,openid
3. every property in component/schema shall define type, and example
4. every api request require to define
    a. with response http status ***200*** and
            - require content type ***application/json***
            - $ref link to suitable schema
    b. ***operationID*** is require to auto generate handle



# Limitation
1. api refer to Schema type "array" , only able to show blank array
2. jwt is not supported yet
3. parameters and apikey's name only support alphanumeric or "_".
4. Response and request body only support application/json


# Technical Info
1. spec.yaml used to generate below objects
    - Allroutes
    - Allschemas
    - Allsecurityschemes
    - Allhandles
    - Allparameters
3. Allroutes generate
    - tell web server each http request (post/get/... ) go to which handle function
    - prepare all handle function, each function sample dataisted in .yaml
    - implement security middleware if define in .yaml 
2. Allschemas use for generate
    - Model (Data type) of each response
    - Example data



Todo:
1. unit test auto run http request using all sample provided in .yaml
2. add configuration for use sample data/use external module





1. change main system, 
    a. if no session automatically go to login microservices
    b. callback url all traffic to microservices
    c. pass environment parameter to microsercices
2.



goal indipendent microservices  for keycloak
1. can redirect
2. support content type strings
3. support different http status response
4. add support for module folder, inside generate dummy methods
    - dummy module
5. fix unit test code and prepare samples




todo:
test
    1. unit test prepare data type mapping response
    2. unit test status mapping
    3. unit test prepare request samples
    4. unit test base on server setting, or base on cli option
    
functions
    easy      
        
        8. refresh token how to know refresh link => define refresh url
        
    medium
        4. application choose yaml's accept mime type
        6. verification of jwt token need more improvement
        6. touch up securityscheme setting not complete use
        8. support examples
    hard
        5. make override option base on analyse file content instead hard code
        5. components support reponses, examples, requestBodies, callbacks
        7. try generator client 
    
    misc:
        auto define port base on yaml file, dont know viable or not
    DONE:   
        to be confirm
            3. docker file
        1. copy api file into /api/api.yaml 
        2. build in swagger-ui, with parameter on-off
        7. user service break oauth into 1 level deeper
        8. use embed to embed resource file
    
    

development
1. pick suitable db driver and query builder
2. prepare db scheme
3. prepare structure of resources
4. add tools
    string process
    db process
    cache process
    


8. 