#!/usr/bin/env bash

#---------test-gateway------------
echo "====test-getEmployer===="
curl -X POST -k https://localhost:50052/example/getEmployer -d '{"employerId": 1}'

echo "====test-getEmployerList===="
curl -X POST -k https://localhost:50052/example/getEmployerList -d '{"employerIndex": 0,"employerCount": 3}'

echo "====test-getEmployerMap===="
curl -X POST -k https://localhost:50052/example/getEmployerMap -d '{}'

echo "====test-getEmployerAll===="
curl -X POST -k https://localhost:50052/example/getEmployerAll -d '{}'

#echo "====test-addEmployerImage===="
#curl -X POST -k https://localhost:50052/example/addEmployerImage -d \
#'{"employerId": 1,"picName": "123test.jpg","picData": "..."}' #picdata 需要载入图片

echo "====test-getEmployerImage===="
curl -X POST -k https://localhost:50052/example/getEmployerImage -d '{"employerId": 3}'


#--------------need-trace----------
#https://localhost:50051/debug/requests
#https://localhost:50051/debug/events
