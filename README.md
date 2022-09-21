# Introduction
### This project still under rapid development, it working but not yet production ready.

The goal of this project is to allow developer setup microservices in shortest time, design with low code concept. 

We know openapi is cool, usually we use swagger or postman for api testing, documentation and design. However, design, develop, testing and documentation is redundant work and spend us lot of time.

To create good rest api, here is the step:
1. Design openapi v3 with swaggerhub, postman or etc openapi GUI design tools, get your spec.yaml (use v3.0 only, v3.1 not yet)
2. Follow steps written at ***How to use*** to generate your microservices program
3. If your .yaml design well, the server will function as expected. Then you can edit your ***openapiserver/openapi/handles.go*** to make your microservices use real data.


 design api standard (spec.yaml), before development. The spec.yaml you did can turn into fully function microservices with step-by-step below. After code generated, just change the function defined in ***openapi/handle.go***, simple and easy.


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
./openapigenerator --apifile="samples/spec.yaml"  --targetfolder="../openapiserverfolder" --projectname="openapiserver" --listen=":9000"
```

3. use your rest api
```bash
cd ../openapiserverfolder
go get .
go build . 
./openapiserver
```

4. Try your rest api http://localhost:9000, to access your mock rest api server.

5. Modify your code  ***openapiserverfolder/openapi/handles.go*** and repeat step 3 to make your rest api function as expected.

# Features
1. Auto prepare data type according schema
2. Auto prepare methods according operationID, response with schema's data type
3. Auto route http traffic to coresponding methods
4. Supported http traffic (GET/POST/PUT/DELETE)
5. Sample work with gin http server
6. Work with application/json response


# Todo
1. basic and bearer jwt authorization
3. use middleware to control access right
3. keep logs
4. split schema by file, add interface for it

# Tips
1. Prepare all schemas sample, and connect to your path
2. Define operationID on every path
3. Define response for http '200' for every http request, and ref to schema
4. Define sample request bodies and response
5. Define your interfaces connect to Model_xxx, and connect ***handles.go*** methods to mdel's method. It can keep your code cleaner.