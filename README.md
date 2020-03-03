# Services

[![Build](https://github.com/MihaiLupoiu/services/workflows/Build/badge.svg?branch=master)](https://github.com/MihaiLupoiu/services/actions)

Repository of services

## Services List

### User

Service for storing user related information

Run Docker Compose:
To run the service, inside `/scripts` there is a docker compose file.
```bash
docker-compose up
```

Build binary: 
```bash
cd src
make build
```

To run tests it requires Postgresql to be started: 
```bash
cd src
make tests
```

Endpoints:

* Create User: `http://localhost:8080/user/add`
Creates a user and does not require authentication
```bash 
curl -s -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/user/add
```

* Login:  `http://localhost:8080/login`
Login using an email an valid password and generates authentication JWT TOKEN. Required in further operations.
```bash 
TOKEN=$(curl -s -X POST -H 'Accept: application/json' -H "Content-Type: application/json" -d "{ \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/login | sed -e 's/^"//' -e 's/"$//' )
```

* GetUser: `http://localhost:8080/user/{id}`
Get user data. If authenticated, it will return user id data, but only his data.
```bash 
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8080/user/1
```

* UpdateUser:  `http://localhost:8080/user/{id}`
Update user data information if authenticated.
```bash 
curl -s -X POST -H "Content-Type: application/json" 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Smith\", \"email\": \"js@gmail.com\",\"password\": \"1234\"}" http://localhost:8080/user/1
```

* DeleteUser:  `http://localhost:8080/user/{id}`
Delete user if authenticated.
```bash 
curl -i -X DELETE -H "Content-Type: application/json" 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8080/user/1
```
