# Services

[![Build](https://github.com/MihaiLupoiu/services/workflows/Build/badge.svg?branch=master)](https://github.com/MihaiLupoiu/services/actions)

Repository of services

## Services List

### User Endpoints 

Service for storing user related information

To run the servin, inside `/scripts` there is a docker compose file. Execute ir like in the nex line and wait a few seconds:
```bash
docker-compose up
```

* Create User: `http://localhost:8080/user/add`
Creates a user and does not require authentication
```bash 
curl -s -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Donals\", \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/user/add
```

* Login:  `http://localhost:8080/login`
Login using an email an valid password and generates autentication JWT TOKEN. Required in furter operations.
```bash 
TOKEN=$(curl -s -X POST -H 'Accept: application/json' -H "Content-Type: application/json" -d "{ \"email\": \"jd@fake.com\",\"password\": \"1234\"}" http://localhost:8080/login | sed -e 's/^"//' -e 's/"$//' )
```

* GetUser: `http://localhost:8080/user/{id}`
Get user data. If autenticated, it will return user id data, but only his data.
```bash 
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8080/user/1
```

* UpdateUser:  `http://localhost:8080/user/{id}`
Update user data information if autenticated.
```bash 
curl -s -X POST -H "Content-Type: application/json" 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" -d "{ \"firstname\": \"Jhon\", \"lastname\": \"Smith\", \"email\": \"js@gmail.com\",\"password\": \"1234\"}" http://localhost:8080/user/1
```

* DeleteUser:  `http://localhost:8080/user/{id}`
Delete user if autenticated.
```bash 
curl -i -X DELETE -H "Content-Type: application/json" 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8080/user/1
```
