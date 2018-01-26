#!/bin/sh

curl -H "Content-Type: application/json" -X PUT -d '{"id":"1","name":"mySoft","version":"1.0"}' http://127.0.0.1:8000/api/add
