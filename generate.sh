#!/bin/bash
# File: generate.sh
# Description: run to generate api

set -e

# Make sure go swagger is installed
echo -e "\033[4mEnsuring swagger generator...\033[0m"
(brew update && brew upgrade go-swagger || true)

# Get latest swagger spec from citadel
#echo -e "\033[4mGetting latest swagger...\033[0m"
#rm -f swagger.yaml
#curl "https://api.fortifi.io/swagger.yaml" --output swagger.yaml

# Generate new models & client from spec
rm -rf client models
echo -e "\033[4mGenerating fortifi API from Swagger spec...\033[0m"
(swagger generate client && go fix ./...)
