package main

import (
  "fmt"
  "net/http"
  "log"
  "github.com/gorilla/mux"
)

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/", IndexHandler)
  router.HandleFunc("/add/{numbers}", AdditionHandler)

  http.Handle("/", router)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello World")
}
