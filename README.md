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

### API
GET /data returns all records where t1.key = t2.t1key
POST /data creates new records in T1 & T2 from body {data: string} (data must not exists)
DELETE /data deletes a record in T1 using the body {key: int} (

