#!/usr/bin/env bash

URL=https://<...>-uc.a.run.app
curl https://<...>-uc.a.run.app
curl $URL/ 

echo curl $URL/stat/mean -X POST -H "Content-Type: application/json" -d '{"data":[1,1,2,3,5]}'

curl $URL/stat/mean -X POST -H "Content-Type: application/json" -d '{"data":[1,1,2,3,5]}'
# > {“data”: [1,1,2,3,5], “result”: 2.4}
echo '{"data":[1,1,2,3,5],"result":2.4}' '<-- expected result'

curl $URL/stat/mean -X PUT -H "Content-Type: application/json" -d '{"data":[1,1,2,3,5]}'
echo 'Method not allowed' '<-- expected result'

curl $URL/stat/mean -X POST -H "Content-Type: application/jso" -d '{"data":[1,1,2,3,5]}'
echo 'Content-Type not allowed' '<-- expected result'

curl $URL/stat/median -X POST -H "Content-Type: text/csv" -d '1,1,2,3,5'
echo '{"data":[1,1,2,3,5],"result": 2}' '<-- expected result'

curl $URL/stat/median -X POST -H "Content-Type: text/csv" -d '1,1,2,3,5,6'
echo '{"data":[1,1,2,3,5,6],"result":2}' '<-- expected result'

curl $URL/stat/median -X PUT -H "Content-Type: text/csv" -d '1,1,2,3,5'
echo 'Method not allowed' '<-- expected result'

curl $URL/stat/median -X POST -H "Content-Type: application/json" -d '1,1,2,3,5'
echo 'Content-Type not allowed' '<-- expected result'
