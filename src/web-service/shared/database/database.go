package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // _ needed
)
//Database info
const (
  host     = "greenexa.cg94kiigpggk.us-west-1.rds.amazonaws.com"
  port     = 5432
  user     = "administrator"
  password = "thisisthepassword"
  dbname   = "ccdata"
  sslmode  = "disable"
)

// DB is exported db
type DB struct {
    *sql.DB
}

//exported for creating a new Postgresql instance
func Connect() (DB, error) {  
    t := "host=%s port=%d user=%s password=%s dbname=%s"
    connectionString := fmt.Sprintf(t, host, port, user, password, dbname)
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal("Invalid DB config:", err)
    }
    
    err = db.Ping()
    if err != nil {
        log.Fatal("DB unreachable:", err)
    }

    return DB{db}, err
}