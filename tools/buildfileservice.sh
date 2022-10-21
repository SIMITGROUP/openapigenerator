#!/bin/bash
rm openapigenerator ; make; ./openapigenerator --apifile="samples/fileservice.yaml"  --targetfolder="../fileservice" --projectname="fileservice" --listen=":9000"  --lang="go" --override="true"
