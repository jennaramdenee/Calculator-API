package main

import (
  "fmt"
  "net/http"
  "log"
  "strings"
  "github.com/gorilla/mux"
  "strconv"
)

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/", IndexHandler)
  router.HandleFunc("/add/{numbers}", AdditionHandler)

  http.Handle("/", router)
  log.Fatal(http.ListenAndServe(":8080", router))
}

func IndexHandler(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
  fmt.Fprintf(w, "Hello World")
}

func AdditionHandler(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
  vars := mux.Vars(r)["numbers"]

  var sum float64 = 0

  s := strings.Split(vars, ",")
  for _, value := range s {
    if num, err := strconv.ParseFloat(value, 64); err == nil {
      sum += num
    }
  }

  fmt.Fprint(w, sum)
}
