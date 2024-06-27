# Receipt Processor Challenge

## Table of Contents
1. [Task](#task)
2. [How to Run](#how-to-run)
3. [Test API ](#test-api)
5. [Next Steps](#next-steps)
### Task
The task is to build a  web service that processes receipts according to the documented API specifications. The service should handle receipt submission, generate a unique ID for each receipt, and calculate points based on predefined rules.
### How to Run
To run the service, follow these steps:

#### 1. Clone the repository:

```bash
git clone https://github.com/sameergopali/Receipt-Processor.git
cd Receipt-Processor
```

#### 2. Install Requirements
You will need the following installed on your local machine
- docker -- [install guide](https://docs.docker.com/get-docker/)
- [go](https://go.dev/doc/install) v1.21.1+ 
-  [swag](https://github.com/swaggo/swag) : This project uses  swag to generate swagger documentation for the apis.
```bash 
go install github.com/swaggo/swag/cmd/swag@latest
```

You can build and run using any of the following options
####  Using Go
- Run tests
```bash
go test ./tests/...
```
- Run swag 
```bash
swag init -d cmd,internal -o docs
```
- Run server
```bash
go run cmd/main.go
```

####  Using Makefile
Ensure that make is installed in the system
```bash
make all
make run
```


#### Using docker:
Make sure docker is running and then run docker compose to start the server
```bash
docker compose up -d
```

### Test API
1. Using curl: 
    - Post Receipt to get id
    ```bash
    curl POST   -H "Content-Type: application/json"   -d '{
      "retailer": "M&M Corner Market",
      "purchaseDate": "2022-03-20",
      "purchaseTime": "14:33",
      "items": [
        {
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        },{
          "shortDescription": "Gatorade",
          "price": "2.25"
        }
      ],
      "total": "9.00"
    }' http://localhost:8080/receipts/process
    ```
    - Get Points Info:
    Use the generated id to get the points information for the receipt
    ```bash
    curl http://localhost:8080/receipts/{id}/points
    ```

2. Use Swagger: The swagger endpoint is available at http://localhost:8080/swagger/index.html. You can test and execute the post and get request in the UI

### Project Details
This project is designed to handle receipt processing in-memory and is only storing the receipt id and points. The Points are calculated based on predefined rules such as retailer name length, total amount characteristics, item descriptions, purchase date and time.

### Next Steps
Some further enhancements that can be done are:
- Adding input validation and error handling for API requests.
- Implementing different logging level and metrics to monitor service performance.
- Adding and extending unit tests to cover additional edge cases as not all unit tests have been added due to time constraints.

