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
5. To execute test we using golang postgres driver
```
go get github.com/lib/pq
```
6. testify 
```
go get github.com/stretchr/testify
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
* Fields mapping via query text & structure tags
* Failure won't occur until runtime
5. Why we don't use database/sql?
* Very fast & straightforward
* Manual mapping SQL fields to variables
* Easy to make mistakes, not caught until runtime
6. 
* To execute test we using golang postgres driver
* Using testify to support run Unit test
7. Implement database transactions in Golang? What is ACID of database transactions.
* Atomicity (A): Either all operations complete successfully or the transaction fails and the db is unchanged
* Consistency (C): The db state must be valid after the transaction. All constraint must be satisfied.
* Isolation (I): Concurrent transaction must not affect each other.
* Durability (D): Data written by a successfully transaction must be recorded in persistent storage 

8. Isolation (I) Read Phenomena
* Dirty Read: A transaction reads data written by other concurrent uncommitted transaction
* Non-repeatable read: A transaction reads the same row twice and sees different value because it has been modified by other committed transaction 
* Phantom read: A transaction re-executes a query to find rows that satisfy a condition and sees different set of rows, due to changes by other committed transaction
* Serialization Anomaly: The result of a group of concurrent committed transactions is impossible to achieve if we try to run them sequentially in any order without overlapping 

9. 4 Standard Isolation Levels
* Low 
* ---> 1. Read Uncommitted (Can see data written by uncommitted transaction)
* ---> 2. Read Committed (Only see data written by committed transaction) 
* ---> 3. Repeatable Read (Same read query always return same result) 
* ---> 4. Serializable (Can achieves same result if execute transactions serially in some order instead of concurrent) 
* ---> High

10. Implement RESTful HTTP API using go gin
```
go get -u github.com/gin-gonic/gin
```

