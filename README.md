# mta-hosting-optimizer


I have created two services first data service which gives us required mock data and second mta_optimizer service which use data service and give less used host_name. 

Deployed Url : http://ec2-35-154-156-239.ap-south-1.compute.amazonaws.com

run command go run . for both the services to run both services and set ENVIRONMENT before run

Project with the name “mta-hosting-optimizer” is hosted on GitLab. GitLab was asking for CI/CD , so I have used AWS for deployment

A major programming language is utilized to build the service. User Golang

A HTTP/REST endpoint to retrieve hostnames having less or equals X active IP addresses exists. PFB endpoint

X is configurable using an environment variable, the default value is set to 1. If env variable is not set value will be 1

IP Configuration data (IpConfig) is provided by a mock service using sample data below. created mock service for this

Code coverage exceeds 90%. 7. Unit & integration tests are present. covered all the functions in unit test

Integrate the test and build phases to Github action Integrated in github action





 
