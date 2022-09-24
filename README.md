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
make
./openapigenerator --apifile="samples/spec.yaml"  --targetfolder="../openapiserverfolder" --projectname="openapiserver" --listen=":9000" --lang="go"
```

3. use your rest api
```bash
cd ../openapiserverfolder
go get .
go build . 
./openapiserver
```

4. Try your rest api http://localhost:9000, to access your mock rest api server. It will run return sample data defined in .yaml file.

5. Put in your real code at below file and repeat (3.)
    a.  ***openapiserverfolder/openapi/routerhandle.go***
    b.  ***openapiserverfolder/openapi/Model_xxx.go***


# Features
1. Auto prepare model/interface according each component's schema
2. Auto prepare path and route to handles according operationID (required), response example data according reference schema and examples
3. Supported http traffic (GET/POST/PUT/DELETE)
4. Build in gin http server
5. Work with application/json response (only)
6. Use middleware control security requirement (support basic and jwt)


# Todo
1. add in missing basic and bearer jwt authorization
2. prepare log system
3. auto generate unit test

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