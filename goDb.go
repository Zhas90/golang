package main

import (
    "database/sql"
    _ "gopkg.in/goracle.v2"
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "log"
)

type Customer struct {
   LAST_NAME  string `json:"lastname"`
	 NAME       string `json:"name"`
	 STATE      string `json:"state"`
	 PAN        string `json:"pan"`
   GUID       string `json:"guid"`
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")

   db, err := sql.Open("goracle", "username/password@localhost:1521/xe")
   if err != nil {
     log.Println(err)
     return
   }
   defer db.Close()

   params := mux.Vars(r)
   query := "SELECT last_name, name, state, pan, guid FROM customer WHERE CUSTOMER_ID = " + params["customerId"]
   var customer Customer
   row := db.Query(query)
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
   r := mux.NewRouter()
   r.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
   log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}