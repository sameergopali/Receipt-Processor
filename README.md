# Receipt Processor Challenge

## Table of Contents
1. [Task](#task)
2. [How to Run](#how-to-run)
3. [Design Questions](#design-question)
5. [Next Steps](#next-steps)
### Task
The task is to build a web service that processes receipts according to the documented API specifications. The service should handle receipt submission, generate a unique ID for each receipt, and calculate points based on predefined rules.
### How to Run
To run the service, follow these steps:

#### 1. Clone the repository:

```bash
git clone https://github.com/sameergopali/Fetch-THA.git
cd Fetch-THA
```

#### 2. Install Requirements
You will need the following installed on your local machine
- docker -- docker [install guide](https://docs.docker.com/get-docker/)
- go [install guide]()
    
#### Using Makefile
```bash
make build
make run
```


#### Using docker:
```bash
docker compose up -d
```

### Test API
Use curl

Use Swagger UI url:

http://localhost:8080/swagger/index.html

### Project Details
This project is designed to handle receipt processing in-memory without data persistence across application restarts. Points are calculated based on predefined rules such as retailer name length, total amount characteristics, item descriptions, purchase date and time.

### Next Steps
For further enhancements, consider:

Adding input validation and error handling for API requests.
Implementing logging and metrics to monitor service performance.
Extending unit tests to cover additional edge cases.
