#!/usr/bin/env bash

BODY='{"duration": 5}'

#BAD REQUEST
#curl -X POST -H "Content-Type: application/json" \
#    -d '{"duration": "5"}' \
#    http://localhost:4000/public/task

curl -w '\nTime: %{time_total} sec \n' -X POST -H "Content-Type: application/json" \
    -d "${BODY}" \
    http://localhost:4000/public/task