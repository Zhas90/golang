package main

import (
    "database/sql"
    _ "gopkg.in/goracle.v2"
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

type Customer struct {
   LAST_NAME  string `json:"lastname"`
	 NAME       string `json:"name"`
	 STATE      string `json:"state"`
	 PAN        string `json:"pan"`
   GUID       string `json:"guid"`
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

   var conf = getConfig()
   db, err := sql.Open("goracle", conf.DbUsername + "/" + conf.DbPassword + conf.DbName)
   if err != nil {
     log.Println(err)
     return
   }
   defer db.Close()

   params := mux.Vars(r)
   query := "SELECT last_name, name, state, pan, guid FROM customer WHERE CUSTOMER_ID = " + params["customerId"]
   var customer Customer
   row := db.QueryRow(query)
   err = row.Scan(&customer.LAST_NAME,
                  &customer.NAME,
                  &customer.STATE,
                  &customer.PAN,
                  &customer.GUID)
   if err != nil {
        log.Println(err)
        return
   }

   json.NewEncoder(w).Encode(customer)
   return
}

func main() {
   var conf = getConfig()
   r := mux.NewRouter()
   r.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
   log.Fatal(http.ListenAndServe(conf.Port, r))
}
