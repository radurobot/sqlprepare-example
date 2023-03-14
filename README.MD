
## Prepare SQL Queries for Fast API Responses with Gin-Gonic

This is an example code that demonstrates how to use prepared SQL queries that aids to speeding up your API responses. This code shows how to execute a simple SQL query to retrieve data from an in-memory SQLite database and return it as JSON.

### How it works

This code uses the Gin-Gonic web framework to handle HTTP requests. When the server receives a GET request for the `/name` endpoint, it reads the `id` parameter from the query string, executes a SQL query to retrieve the corresponding data from the SQLite database, and returns the result as JSON.

The SQL query is prepared in advance using the `db.Prepare` function, which compiles the query into a reusable statement that can be executed multiple times. The `stmt.QueryRow` function is then used to execute the prepared statement with the given `id` parameter, and the result is scanned into a variable using the `Scan` function.

### How to run

To run this code, you'll need to have Go installed on your machine. You can then download or clone the repository and run the following command:

`go run main.go` 

This will start the server on port 8081. You can then send a GET request to `http://localhost:8081/name?id=1` to retrieve the data for the `id=1` record from the database.

### Note
For a detailed explanation of how to use SQL queries to speed up your API responses with Gin-Gonic, please refer to this article: [Speed Up Your API by preparing SQL Queries in Go](https://medium.com/@ifrimradumdc/speed-up-your-api-by-preparing-sql-queries-in-go-ddfaa2c27c41)