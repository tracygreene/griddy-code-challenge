package model

import (
    "database/sql"
    "time"
    "fmt"
    "web-service/shared/database"
)


//record defines model for t1 & t2 join
type record struct {
  T1_Key int 
  T1_Data string
  T2_Key int
  T2_T1key int
  T2_Data string
  T2_CreatedOn time.Time
}

type records struct {
  Records []record
}

// GetAll gets all the records from join of t1 & t2
func GetAll() ([]record, error) {
    //connect to database
    DB, err := database.Connect()
    
    sqlStatement := `
    SELECT * 
    FROM t1
    INNER JOIN t2
    ON t1.key = t2.t1key
    ORDER BY t1.key ASC`
    
    //get rows where t1.key = t2.t1key
    rows, err := DB.Query(sqlStatement)
    if err != nil  {
       return nil, err
    }
        
    defer rows.Close()
        
    results := records{}
        
    for rows.Next() {
        result := record{}
        err = rows.Scan(
          &result.T1_Key,
          &result.T1_Data,
          &result.T2_Key,
          &result.T2_T1key,
          &result.T2_Data,
          &result.T2_CreatedOn,
        )
        if err != nil {
           return nil, err
        }
        
        results.Records = append(results.Records, result)
    }
    err = rows.Err()
    if err != nil {
       return nil, err
    }
        
    DB.Close()

    return results.Records, nil
}

// Post inserts all the specific record into the database
func Post(data string) (newKey int) {

    //connect to database
    DB, err := database.Connect()
        
    sqlStatement := `
    INSERT INTO t1 (data)
    VALUES ($1)
    ON CONFLICT (data)
    DO NOTHING
    RETURNING key`
        
    key := 0
        
    //insert row into t1 table
    err = DB.QueryRow(sqlStatement, data).Scan(&key)
    if err != nil && err != sql.ErrNoRows {
         panic(err)
    }
        	
    //create t2 string made up of the data posted & the new key from the inserted t1 record
    t2Str := fmt.Sprintf("%s %d", data, key)
    
    t2Key := 0
    	    
    //insert into t2 
    sqlStatement = `
    INSERT INTO t2 (t1key, data)
    VALUES ($1, $2)
    RETURNING key`
    
    //insert into t2 
    err = DB.QueryRow(sqlStatement, key, t2Str).Scan(&t2Key)
    if err != nil {
    	panic(err)
    } 
        
    fmt.Println("New key is:", key)
    
    //close database connection
    DB.Close()

    return key
 }
  
// Delete removes a record from the database
func Delete(deleteKey int) (rowCount int64) {
    DB, err := database.Connect()
          
    sqlStatement := `
    DELETE FROM t1
    WHERE key = $1`
          
    //delete record from t1 (the associated t2 record is auto-deleted) 
    res, err := DB.Exec(sqlStatement, deleteKey)
    if err != nil {
       panic(err)
    }
          
    count, err := res.RowsAffected()
    if err != nil {
       panic(err)
    }
          	
    //close database connection
    defer DB.Close()
  
    return count
}