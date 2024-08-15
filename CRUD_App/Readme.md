
- Project setup instructions 

1. packages:
go mod init Crud_Api  
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/gorilla/mux 
go get github.com/joho/godotenv 
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/http-swagger 


2. create the Database in pgadmin  db name is : CRUD-API

3. run cmd : go run main.go 

- API endpoint documentation.

User Api :

Post Request: http://localhost:8000/api/v1/users
Get Request:http://localhost:8000/api/v1/users
Put Request:http://localhost:8000/api/v1/users/1  
Delete Request: http://localhost:8000/api/v1/users/2

 - How to run the tests.
 I created the user_test.go file in the controllers folder : controller/user_test

 - Deployment instructions.

 - Use a tool like Swagger to generate API documentation
 I created the swagger documents file in the docs folder 