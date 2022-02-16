# KV-Store

**KV-Store** is a REST-API service that works as an in memory key-value store.

[![Build Docker Image](https://github.com/semihsemih/kv-store/actions/workflows/container.yml/badge.svg)](https://github.com/semihsemih/kv-store/actions/workflows/container.yml)
[![Check Lint with golangci-lint](https://github.com/semihsemih/kv-store/actions/workflows/lint.yml/badge.svg)](https://github.com/semihsemih/kv-store/actions/workflows/lint.yml)
[![Run Tests](https://github.com/semihsemih/kv-store/actions/workflows/test.yml/badge.svg)](https://github.com/semihsemih/kv-store/actions/workflows/test.yml)

## Installation
You need require **Docker** and **Git**

1. Clone this repo.
```
git clone https://github.com/semihsemih/kv-store
```

2. Run the docker-compose.yml file and start the server
```
docker-compose up -d
```

## Usage
### Set new key-value
```
/set?key=foo&value=bar
```
### Get value given by key
```
/get?key=foo
```
### Flush all storage
```
/flush
```

## Logging and Data Storage
When you start the application for the first time, a folder named ``store`` is created. Folder contains server.log file and data.json file. 
- The server.log file is the log file that logging incoming requests to the server.
- The data.json file is the file that stores the key-values and data is read when the system is turned on.
