/**********************************************************
 TG 9/25/2018: Modified to use database instead of variable
**********************************************************/
package main

import (
    "log"
    "net/http"
   
    "web-service/controller"
    
    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
)


func main() {
    router := mux.NewRouter()
    
    router.HandleFunc("/data", controller.GetDataEndpoint).Methods("GET")
    router.HandleFunc("/data", controller.CreateDataEndpoint).Methods("POST")
    router.HandleFunc("/data", controller.DeleteDataEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":80", router))
}
