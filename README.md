# GraphQL DEMO

Demo application for learning GraphQL.

## Getting Started

ref. https://bit.ly/2PCme9M


### Prerequisites

#### Node
- v8.12.0
- (npm version 6.4.1)

#### Golang
- v1.11.1
- https://github.com/kardianos/govendor
- https://github.com/jessevdk/go-assets-builder

## Setup

### 1. Setup MySQL
```bash
# create database
$ mysql -h172.0.0.1 -uusername -p < db/create_database.sql

# migration
$ mysql -h172.0.0.1 -uusername -p graphql-demo < db/create_table.sql

# insert seed data
$ mysql -h172.0.0.1 -uusername -p graphql-demo < db/seed.sql
```

### 2. Set Environment Variable
```bash
# data source
$ export dsn = "username:password@tcp(170.0.0.1)/graphql-demo"

# port number
$ export PORT = "3000"

```

### 3. Install Dependent Libraries
```bash
# install npm packages
$ npm install

# install golang packages
$ govendor sync
```

### 4. Launch Application
```bash
$ make up
```

### 5. Go Application!
|  page  |  url  |
| ---- | ---- |
|  application  |  http://localhost:3000  |
|  graphiql  |  http://localhost:3000/graphql  |
