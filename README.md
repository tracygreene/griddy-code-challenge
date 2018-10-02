# griddy-code-challenge
A Golang based web-service

### Tech Stack
- Virtual server: Amazon EC2   
- Database: Amazon RDS PostgreSQL
- Back-end: Golang
- API testing: Postman

### Database Schema
T1 (PK int key, unique string value)

T2 (PK int key, FK int t1key, string value, default date)

### Project Structure
- web-service
  - controller
    - controller.go - handlers
  - model
    - repository.go - defines data database model and executes CRUD operations against database
  - shared
    - database
      - database.go - database connection
  - main.go - main package to run service and defines routers
  

### CRUD of data
GET /data returns all records where t1.key = t2.t1key

POST /data creates new records in T1 & T2 from body json: 
```
{"data": string}
```
(data must not exists)

DELETE /data deletes a record in T1 using the body json: 
```
{"key": int} 
```
(the referenced record in T2 is auto-deleted)

### Test via curl
```curl
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"data":"string goes here"}' \
  http://ec2-13-57-181-186.us-west-1.compute.amazonaws.com:80/data
```

```
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://ec2-13-57-181-186.us-west-1.compute.amazonaws.com:80/data
```

