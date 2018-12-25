package main

import (
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "log"
    "os"
)

type Configuration struct {
   Port       string
   DbName     string
   DbUsername string
   DbPassword string
}

func getConfig() Configuration {
  //reading config params from file
  file, err := os.Open("config.json")
  if err != nil {
    log.Fatal(err)
  }
  decoder := json.NewDecoder(file)
  var config Configuration
  err = decoder.Decode(&config)
  if err != nil {
    log.Fatal(err)
  }

  return config
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")
   params := mux.Vars(r)
   customerId := params["customerId"]

   json.NewEncoder(w).Encode(getCustomerFromDb(customerId))
   return
}

func getSession(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")
   params := mux.Vars(r)
   sessionId := params["sessionId"]

   json.NewEncoder(w).Encode(getSessionFromDb(sessionId))
   return
}

func main() {
   var conf = getConfig()
   r := mux.NewRouter()
   r.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
   r.HandleFunc("/sessions/{sessionId}", getSession).Methods("GET")
   log.Fatal(http.ListenAndServe(conf.Port, r))
}
