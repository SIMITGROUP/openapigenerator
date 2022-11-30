# Openapi Generator
An openapiv3 generator which can generate micro-services for GO Language

## Content
<hr/>

- [Openapi Generator](#gin-web-framework)
  - [Contents](#contents)
  - [Introduction](#introduction)
  - [Quick Start](#quickstart)
  - [Run Docker](#rundocker)



## Introduction
We know openapi is cool, usually we use swagger or postman for api testing, documentation and design. However, design, develop, testing and documentation is redundant work and spend us lot of time.

This generator able to generate micro-services with below capability:
1. working gin mock server, you can modified it to become real microservice
2. unit test for all restapi
3. build in swagger-ui to allow you test api easily
4. easily generate docker image
5. build in support security scheme


## Quickstart
This step guide you generate micro-service source code and open it with vscode:
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



## Run Docker
Project come with few command help you run microservice in docker environment. Check `Makefile` to know more detail

1. create docker container in devnetwork.
```
cd myproject
make dockerremove ; make docker;  make dockerrun #run docker, under devnetwork
make dockershell # access docker shell 
```
2 You may edit `Makefile` if you wish to pass create docker container with more useful setting



