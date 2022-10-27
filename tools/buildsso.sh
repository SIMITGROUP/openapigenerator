#!/bin/bash
./openapigenerator --apifile="samples/sso.yaml"  --targetfolder="../userservice" --projectname="userservice" --listen=":9000"  --lang="go" --override="true"
