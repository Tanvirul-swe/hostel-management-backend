# Hostel-Management-Backend

This project covers a completely managed hostel. Especially the University hostel. 

## Usage
At first, you have to clone the project in your desired directory. Then Open the Project in your suitable IDE.
`Note: Must be pre-installed GO Language latest version`

## Development Environment setup
Now update `.env` configuration. 
If haven't a PostgreSQL database credential. Can use https://www.elephantsql.com to provide free Database storage.

```env
    DB_URL = Database URL
    DB_USERNAME = Database User Name
    DB_NAME = Database Name
    PORT = Local host machine Port
    HOST = Database Host
    PASSWORD = Database Password
    SSL_MODE = "disable"
    JWT_SECRET = JWT KEY
```

## Initialize
Here few steps for initializing the project's required package. Open the terminal and hit the command
`Note: Make Sure you are in the current project directory

Here Is the Example : 

![Screenshot from 2023-09-15 21-54-50](https://github.com/Tanvirul-swe/flutter_blue_serial/assets/75753499/96ed98fa-3de2-47c8-aa6f-b2c75e87df70)

 **First Command**
`go mod tidy`
 **Second Command**
`go get`

## Run
Open the terminal and hit `go run main.go`

## Go Versions

- GO 1.6: >= 1.21

## Maintainers

- [Md. Tanvirul Islam](https://github.com/Tanvirul-swe)

