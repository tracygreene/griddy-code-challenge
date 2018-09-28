package controller

import (
    "encoding/json"
    "fmt"
    "net/http"

    "web-service/model"
)

//Get HANDLER
func GetDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    results, err := model.GetAll()
    
    //convert to json
    out, err := json.Marshal(results)
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

//data from post
type postData struct {
  Data string
}

//post HANDLER 
func CreateDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    //get the body
    decoder := json.NewDecoder(r.Body)
    
    var post postData
    json_err := decoder.Decode(&post)
    if json_err != nil {
    	panic(json_err)
    }
    
    data := post.Data
    
    response := " "
    
    key := model.Post(data)
    //if the data already exists then row isn't inserted
    if key < 1 {
    	response = fmt.Sprintf("%s already exists, no duplicates allowed", data)
    } else {
    	response = fmt.Sprintf("%s was successfully added with the key %d", data, key)
    }
    
    json.NewEncoder(w).Encode(response)
    
}

//data from delete
type deleteData struct {
  Key int
}

//Delete HANDLER
func DeleteDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    //get the body
    decoder := json.NewDecoder(r.Body)
        
    var key deleteData
    
    json_err := decoder.Decode(&key)
    if json_err != nil {
        panic(json_err)
    }
        
    deleteKey := key.Key
    
    count := model.Delete(deleteKey)
        
    response := " "
        
    //if no rows where returned, nothing was deleted
    if count < 1 {
        response = fmt.Sprintf("no record with key %d was found", deleteKey)
    } else {
    	response = fmt.Sprintf("the record with key %d was successfully deleted", deleteKey)
    }
        
    json.NewEncoder(w).Encode(response)  
}

