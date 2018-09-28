/**********************************************************
 TG 9/25/2018: Modified to use database instead of variable
**********************************************************/
package main

import (
    //"encoding/json"
    "log"
    "net/http"
    //"database/sql"
    //"fmt"
    //"time"
    "web-service/controller"
    
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)

//Database info
/*const (
  host     = "greenexa.cg94kiigpggk.us-west-1.rds.amazonaws.com"
  port     = 5432
  user     = "administrator"
  password = "thisisthepassword"
  dbname   = "ccdata"
  sslmode  = "disable"
)

//data from delete
type deleteData struct {
  Key int
}

//record defines model for t1 & t2 in database
/*type record struct {
  T1_Key int
  T1_Data string
  T2_Key int
  T2_T1key int
  T2_Data string
  T2_CreatedOn time.Time
}*/

/*type records struct {
  Records []record
}

//Get HANDLER
func GetDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    //connect to database
    /*db := connect()

    sqlStatement := `
    SELECT * 
    FROM t1
    INNER JOIN t2
    ON t1.key = t2.t1key
    ORDER BY t1.key ASC`

    //get rows where t1.key = t2.t1key
    rows, err := db.Query(sqlStatement)

    if err != nil  {
    	panic(err)
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
    		panic(err)
    	}
    
    	results.Records = append(results.Records, result)
    }
    
    err = rows.Err()
    if err != nil {
    	panic(err)
    }
    
    //close database connection
    defer rows.Close()*/
    //results, err := repository.GetAll())
    
    //convert to json
    /*out, err := json.Marshal(results)
    if err != nil {
    	http.Error(w, err.Error(), 500)
    	return
    }
    
    //Set Content-Type header so that clients will know how to read response
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    //Write json response back to response 
    w.Write(out)
    
}

//MODEL for post
type postData struct {
  Data string
}

//HANDLER for post
func CreateDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    //get the body
    decoder := json.NewDecoder(r.Body)
    
    var post postData
    json_err := decoder.Decode(&post)
    if json_err != nil {
    	panic(json_err)
    }
    
    data := post.Data
    
    //connect to database
    db := connect()
    
    sqlStatement := `
    INSERT INTO t1 (data)
    VALUES ($1)
    ON CONFLICT (data)
    DO NOTHING
    RETURNING key`
    
    key := 0
    
    //insert row into t1 table
    err := db.QueryRow(sqlStatement, data).Scan(&key)
    
    if err != nil && err != sql.ErrNoRows {
        panic(err)
    }
    
    response := " "
    
    //if the data already exists then row isn't inserted
    if err == sql.ErrNoRows {
    	response = fmt.Sprintf("%s already exists, no duplicates allowed", data)
    } else {
    	response = fmt.Sprintf("%s was successfully added with the key %d", data, key)
    	
    	//create t2 string
	t2Str := fmt.Sprintf("%s %d", data, key)
	    
	//insert string into t2 
	sqlStatement = `
	INSERT INTO t2 (t1key, data)
	VALUES ($1, $2)
	RETURNING key`
	        
	err = db.QueryRow(sqlStatement, key, t2Str).Scan(&key)
	        
	if err != nil {
		panic(err)
    	} 
    }
    
    //close database connection
    defer db.Close()
    
    json.NewEncoder(w).Encode(response)
    
}

//Delete method
func DeleteDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    //get the body
    decoder := json.NewDecoder(r.Body)
        
    var key deleteData
    
    json_err := decoder.Decode(&key)
    if json_err != nil {
        panic(json_err)
    }
        
    deleteKey := key.Key
    
    //connect to database
    db := connect()
    
    sqlStatement := `
    DELETE FROM t1
    WHERE key = $1`
    
    //delete record from t1 (the associated t2 record is auto-deleted) 
    res, err := db.Exec(sqlStatement, deleteKey)
    if err != nil {
      panic(err)
    }
    
    count, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
        
    response := " "
        
    //if no rows where returned, nothing was deleted
    if count < 1 {
        response = fmt.Sprintf("no record with %d was found", deleteKey)
    } else {
    	response = fmt.Sprintf("the record with key %d was successfully deleted", deleteKey)
    }
    	
    //close database connection
    defer db.Close()
        
    json.NewEncoder(w).Encode(response)
    
}

func connect() *sql.DB {  
    t := "host=%s port=%d user=%s password=%s dbname=%s"
    connectionString := fmt.Sprintf(t, host, port, user, password, dbname)
    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        fmt.Println("NOT OPEN")
        panic(err)
    }
    
    err = db.Ping()
    if err != nil {
    	fmt.Println("NOT PING")
        panic(err)
    }

    return db
}*/

func main() {
    router := mux.NewRouter()
    
    router.HandleFunc("/data", controller.GetDataEndpoint).Methods("GET")
    router.HandleFunc("/data", controller.CreateDataEndpoint).Methods("POST")
    router.HandleFunc("/data", controller.DeleteDataEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}