# Golang Backend Server Experiment for wodss Module
## Description
Go Backend Server with MySQL DB, for Study Planer from wodss module

## Installation
### Assuming Go is installed in the version 1.13.8 or newer
* Clone Repo
* Setup MySQL DB "wodss" other used parameters can be found in the .env file
* Remove or Comment HTTPS related Code in main.go, because the cretificats are on our server
* Run it with "go run main.go"

 # Testing
 * Import Test Suites in Postman
 * Change PRODUCTION to false in .env file, drops all tables and creates them again on server start
 * Setup environment with url-variable: http://localhost:3000/api for local testing
 * Run Suites on empty DB
