# Pre
1. Install golang-migrate with command 
```
brew install golang-migrate
```
2. Install docker to setup database locally

3. Install SQLC 
```
brew install sqlc
```
4. After install sqlc successfully, go to root folder and run 
```
sqlc init
```
# DB 
1. Using postgresql
2. Using SQLC
* Very fast & easy to use
* Automatic code generation
* Catch SQL query errors before generating codes
* Full support Postgres. MySQL is experimental
3. Why we don't use GORM? 
* CRUD functions already implemented, very short production code
* Must learn to write queries using GORM's function
* Run slowly on high load
4. Why we don't use SQLX?
* Quite fast & easy to use
* Fields mapping via query text & struct tags
* Falure won't occur until runtime
5. Why we don't use database/sql?
* Very fast & straightforward
* Manual mapping SQL fields to variables
* Easy to make mistakes, not caught until runtime

