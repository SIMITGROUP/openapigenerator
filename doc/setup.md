
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
