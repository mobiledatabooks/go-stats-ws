#!/usr/bin/env bash

# manual testing:
# start the server running 
# go run main.go

curl http://localhost:8085/stat/mean -X POST -H "Content-Type: application/json" -d '{"data":[1,1,2,3,5]}'
# > {“data”: [1,1,2,3,5], “result”: 2.4}
curl http://localhost:8085/stat/mean -X PUT -H "Content-Type: application/json" -d '{"data":[1,1,2,3,5]}'
# > Method not allowed
curl http://localhost:8085/stat/mean -X POST -H "Content-Type: application/jso" -d '{"data":[1,1,2,3,5]}'
# > Content-Type not allowed

curl http://localhost:8085/stat/median -X POST -H "Content-Type: text/csv" -d '1,1,2,3,5'
#   {"data": [1 1 2 3 5], "result": 2} 
curl http://localhost:8085/stat/median -X POST -H "Content-Type: text/csv" -d '1,1,2,3,5,6'
# {"data":[1,1,2,3,5,6],"result":2}
curl http://localhost:8085/stat/median -X PUT -H "Content-Type: text/csv" -d '1,1,2,3,5'
# Method not allowed
curl http://localhost:8085/stat/median -X POST -H "Content-Type: application/json" -d '1,1,2,3,5'
# Content-Type not allowed