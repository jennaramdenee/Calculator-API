package main

import (
  "fmt"
  "net/http"
  "log"
  "strings"
  "github.com/gorilla/mux"
  "strconv"
)

type Answer struct {
  Answer  float64
}

func main(){
  router := mux.NewRouter()
  router.HandleFunc("/", IndexHandler)
  router.HandleFunc("/add/{numbers}", AdditionHandler)
  router.HandleFunc("/multiply/{numbers}", MultiplicationHandler)

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

  s := strings.Split(vars, ",")

  var sum float64 = 0
  for _, value := range s {
    if num, err := strconv.ParseFloat(value, 64); err == nil {
      sum += num
    }
  }

  fmt.Fprint(w, sum)
}

func MultiplicationHandler(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
  vars := mux.Vars(r)["numbers"]

  s := strings.Split(vars, ",")

  var sum float64 = 1
  for _, value := range s {
    if num, err := strconv.ParseFloat(value, 64); err == nil {
      sum *= num
    }
  }

  fmt.Fprint(w, sum)

}
