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
  inputs := GetURLInput(r, "numbers")
  sum := Reduce(float64(0), inputs, Add)
  fmt.Fprint(w, sum)
}

func MultiplicationHandler(w http.ResponseWriter, r *http.Request){
  defer r.Body.Close()
  inputs := GetURLInput(r, "numbers")
  sum := Reduce(float64(1), inputs, Multiply)
  fmt.Fprint(w, sum)
}

func Multiply(x, y float64) float64 {
  return x * y
}

func Add(x, y float64) float64 {
  return x + y
}

func GetURLInput(r *http.Request, s string) []string {
  vars := mux.Vars(r)[s]
  return strings.Split(vars, ",")
}

func Reduce(output float64, inputs []string, fn func(float64, float64) float64) float64 {
  for _, value := range inputs {
    if num, err := strconv.ParseFloat(value, 64); err == nil {
      output = fn(output, num)
    }
  }
  return output
}
