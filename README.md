##Gin-MVC-MySql-API
Implement Go with gin gonic, mvc structure and mysql as API project

###Installation
1. Create database `pos_db` and the table `products`
    ```mysql
    CREATE TABLE `pos_db`.`products` (
      `id` INT NOT NULL AUTO_INCREMENT,
      `name` VARCHAR(255) NULL,
      `code` VARCHAR(255) NULL,
      `sell_price` DECIMAL(50,6) NULL,
      `buy_price` DECIMAL(50,6) NULL,
      `created_at` DATETIME NULL,
      `updated_at` DATETIME NULL,
      `deleted_at` DATETIME NULL,
      `product_bundle` TINYINT(1) DEFAULT '1',
      PRIMARY KEY (`id`));
    ```
  
2. Make sure you already installed Go and some libraries bellow just run these command:
    ```shell script
    > go get -u github.com/gin-gonic/gin
    > go get github.com/shopspring/decimal
    > go get -u github.com/go-sql-driver/mysql
    ```

3. Clone this Repository on your Go-Workspace. In this case I clone on folder `Documents/GoWorkSpace/src/github.com/mfirmanakbar/gin-mvc-mysql-api/`. As we know, if we're start to make a golang project we should put it into the Go Workspace.

4. Configuration, in this case i'm using GoLand as IDEA
    - Add new environment and choose a template Go Build
    - Package Path is `github.com/mfirmanakbar/gin-mvc-mysql-api` with these environment:
        ```shell script
        mysql_users_username=YOUR_DATABASE_USER
        mysql_users_password=YOUR_DATABASE_PASSWORD
        mysql_users_host=127.0.0.1:3306
        mysql_users_schema=pos_db
        ```

5. Then Run it

###API Documentation
1. Create Product
    - Method `POST`
    - URL `http://localhost:8282/products`
    - Header `go-access-token`, the value of this header is optional, we can fill it by empty or fake header value
    - Body Request
        ```json
        {
            "name": "Kopiko",
            "code": "COD001",
            "sell_price": 50000.0,
            "buy_price": 50000.0,
            "product_bundle": 0
        }
        ```
    - Response if the header `go-access-token` empty
        ```json
        {
            "name": "Excelso",
            "code": "COD002",
            "price": "50000",
            "product_bundle": 0
        }      
        ```
    - Response if the header `go-access-token` not empty
        ```json
        {
             "id": 1,
             "name": "Kopiko",
             "code": "COD001",
             "sell_price": "50000",
             "buy_price": "50000",
             "created_at": "2020-01-19 09:03:35",
             "updated_at": "",
             "deleted_at": "",
             "product_bundle": 0
        }
        ```
    