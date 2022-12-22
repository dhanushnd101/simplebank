# Simple Bank (Dhanush)

This is the code writtern after taken the course by [TECH SCHOOL](https://dev.to/techschoolguru) on  [Youtube channel](https://www.youtube.com/c/TECHSCHOOLGURU).

![Backend master class](/backend-masterclass.jpg)

The project had different phases such as design, develop and deploy a backend web service from scratch



## Course

### Section 1: [Working with database Postgres]

- 1: Design DB schema and generate SQL code with dbdiagram.io
- 2: Install & use Docker + Postgres + TablePlus to create DB schema
- 3: How to write & run database migration in Golang
- 4: Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc
- 5: Write unit tests for database CRUD with random data in Golang
- 6: A clean way to implement database transaction in Golang
- 7: DB transaction lock & How to handle deadlock in Golang
- 8: How to avoid deadlock in DB transaction? Queries order matters!
- 9: Deeply understand transaction isolation levels & read phenomena in MySQL & 
- 10: Setup Github Actions for Golang + Postgres to run automated tests

### Section 2: [Building RESTful HTTP JSON API Gin]

- 11: Implement RESTful HTTP API in Go using Gin
- 12: Load config from file & environment variables in Go with Viper
- 13: Mock DB for testing HTTP API in Go and achieve 100% coverage
- 14: Implement transfer money API with a custom params validator
- 15: Add users table with unique & foreign key constraints in PostgreSQL
- 16: How to handle DB errors in Golang correctly
- 17: How to securely store passwords? Hash password in Go with Bcrypt
- 18: How to write stronger unit tests with a custom gomock matcher
- 19: Why PASETO is better than JWT for token-based authentication?
- 20: How to create and verify JWT & PASETO token in Golang
- 21: Implement login user API that returns PASETO or JWT access token in Go
- 22: Implement authentication middleware and authorization rules in Golang using Gin

### Section 3: [Deploying the application to production Kubernetes + AWS]

- 23: Build a minimal Golang Docker image with a multistage Dockerfile
- 24: How to use docker network to connect 2 stand-alone containers
- 25: How to write docker-compose file and control service start-up orders with wait-for
- 26: How to create a free tier AWS account
- 27: Auto build & push docker image to AWS ECR with Github Actions
- 28: How to create a production DB on AWS RDS
- 29: Store & retrieve production secrets with AWS secrets manager
- 30: Kubernetes architecture & How to create an EKS cluster on AWS
- 31: How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS
- 32: How to deploy a web app to Kubernetes cluster on AWS EKS
- 33: Register a domain name & set up A-record using Route53
- 34: How to use Ingress to route traffics to different services in Kubernetes
- 35: Automatic issue TLS certificates in Kubernetes with Let's Encrypt
- 36: Automatic deploy to Kubernetes with Github Action

## Simple bank service

The service that was created will provide APIs for the frontend to do following:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.


### Setup infrastructure


You can get started by calling the following commands


- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```

### How to generate code

- Generate schema SQL file with DBML:

    ```bash
    make db_schema
    ```

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```

## Deploy to kubernetes cluster

- Run the deply scripts to run the application on cloud EKS

```bash
kubectl apply -f eks/aws-auth.yaml
kubectl apply -f eks/deployment.yaml
kubectl apply -f eks/service.yaml
```