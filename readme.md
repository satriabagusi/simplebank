<!--
Author: Satria Bagus(satria.bagus18@gmail.com)
readme.md (c) 2023
Desc: description
Created:  2023-06-30T13:24:05.153Z
Modified: !date!
-->


# Simple Bank API using GO 

![Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1280px-Go_Logo_Blue.svg.png)



## Installation
1. Firstly, Clone or Download this Repo to your local directory
2. You need GO (https://go.dev/doc/install)
3. Open terminal from your IDE or direct open terminal from local directory  

    
## Run Locally

Go to the project directory

```bash
  cd simplebank
```

Run the Project from Terminal

```bash
  go run cmd/main.go
```

**API Documentation: **

```http
POST /api/v1/login
```
```http
POST /api/v1/logout
```
```http
POST /api/v1/transaction/create
```

**Data Example: **
```json
{
    "username": "john",
    "password": "john123"
}
```
OR

```json
{
    "username": "doe",
    "password": "doe123"
}
```

**Using Endpoint Transaction & Logout: **

Put your token and used as Bearer token

Example :
```sh
Bearer token : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEiLCJ1c2VybmFtZSI6ImpvaG4iLCJlbWFpbCI6ImpvaG5AbWFpbC5jb20iLCJleHAiOjE2ODgxMzc5MjAsImlhdCI6MTY4ODEzMDcyMH0.3b39fFQ6D7SjrrjTZYE9mROm_qX1pYAhwcG0FGplSmc
```
