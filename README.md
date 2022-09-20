# Introduction
### This project still under P.O.C stage, not yet production ready.

The goal of this project is to allow us develop microservices in shortest time, design with low code concept. 

We know openapi is cool, usually we use swagger or postman for api testing, documentation and design. However, design, develop, testing and documentation is redundant work and spend us lot of time.

In proper API design, we shall design api standard (spec.yaml), before development. The spec.yaml you did can turn into fully function microservices with step-by-step below. After code generated, just change the function defined in ***openapi/userfunction.go***, simple and easy.


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
./openapigenerator --apifile="samples/spec.yaml"  --targetfolder="../openapiserverfolder" --projectname="openapiserver"
```

3. use your rest api
```bash
cd ../openapiserverfolder
go get .
go build . 
./openapiserver
```

4. Try your api according your spec

5. 

# Features
1. Auto prepare data type according schema
2. Auto prepare methods according operationID, response with schema's data type
3. Auto route http traffic to coresponding methods
4. Supported http traffic (GET/POST/PUT/DELETE)
5. Sample work with gin http server
6. Work with application/json response


# Todo
1. try connect database
2. try basic and bearer jwt authorization
3. try control path authentication
4. try connect program connect outside userfunction.go
5. logs
6. openapi server environment variables, and flag, arguments
7. api generator support flat, arguments