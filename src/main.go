package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

//Store data from post method
var dataStr = " "

//Get method
func GetDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    json.NewEncoder(w).Encode(dataStr)
    
}

//Post method
func CreateDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    params := mux.Vars(r)
    dataStr = params["str"]
    json.NewEncoder(w).Encode(dataStr)
    
}

//Delete method
func DeleteDataEndpoint(w http.ResponseWriter, r *http.Request) {
    
    dataStr = ""
    json.NewEncoder(w).Encode(dataStr)
}

func main() {
    router := mux.NewRouter()
    
    router.HandleFunc("/data", GetDataEndpoint).Methods("GET")
    router.HandleFunc("/data/{str}", CreateDataEndpoint).Methods("POST")
    router.HandleFunc("/data", DeleteDataEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}