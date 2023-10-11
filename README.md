# GO HTTPRouter CRUD
Simple CRUD with golang HTTPRouter

## **Packages used**
- github.com/joho/godotenv/cmd/godotenv@latest
- github.com/go-sql-driver/mysql
- github.com/julienschmidt/httprouter
- github.com/go-playground/validator/v10
- github.com/stretchr/testify

## **Run the migration**
To create a main database and testing database with migrations, please use <a href="https://github.com/golang-migrate/migrate">golang-migrate</a>
```
migrate -database "mysql://user:password@tcp(host:port)/dbname?query" -path migrations up
migrate -database "mysql://user:password@tcp(host:port)/testdbname?query" -path migrations up
```

## **Structure**
Based on repository pattern, this project use:
- Repository layer: For accessing db in the behalf of project to store/update/delete data
- Service layer: Contains set of logic/action needed to process data/orchestrate those data
- Models layer: Contains set of entity/actual data attribute
- Controller layer: Acts to mapping users input/request and presented it back to user as relevant responses

## **API Endpoints**
You can use <a href="https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi">this VSCode Extension</a> to preview the OpenApi .yml file

