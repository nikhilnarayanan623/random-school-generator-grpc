## To Run The Project

## 1 Clone The Repo And Checkout To Project
```
git clone https://github.com/nikhiloayaw/generate-random-school-with-grpc
```
```
cd ./generate-random-school-with-grpc
```

## 2 Set Up and Run School Service
##### 1 Checkout to School service Directory
```
cd ./school-service
```
##### 2 Download the dependencies
```
go mod tidy
```
##### 3 Setup Env
create .env file and add the below details
```.env
## example
 SCHOOL_SERVICE_HOST=localhost
 SCHOOL_SERVICE_PORT=50051
```
##### 4 Run The School Service
```
go run cmd/api/main.go
```

## 2 Set Up and Run API Gateway Service
##### Take A New Terminal

##### 1 Checkout to API Gateway service Directory
```
cd ./api-gateway
```
##### 2 Download the dependencies
```
go mod tidy
```

##### 3 Setup Env
create .env file and add the below details
```.env
## example
  API_PORT=8000
 SCHOOL_SERVICE_HOST=localhost
 SCHOOL_SERVICE_PORT=50051
```
##### 4 Run The API Gateway service
```
go run cmd/api/main.go
```

## Available API( If port is update then change the port as well)
1. ###CREATE SCHOOL
#### Url:  (http://localhost:8000/school)
#### Params(query):  name="name of the school"
#### Header:  ```"Content-Type" = "application/json" ``` for output in json
#### **If content type not provided out will be in ```Excel```**
   
