#!/usr/bin/env bash

docker run --rm -p 81:8080 -e API_URL=http://127.0.0.1:8081/swagger.json swaggerapi/swagger-ui
