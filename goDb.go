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
	 NAME      string `json:"name"`
	 STATE     string `json:"state"`
	 PAN       string `json:"pan"`
   GUID      string `json:"guid"`
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
   customerId := params["customerId"]
   rows, err := db.Query("SELECT name, state, pan, guid FROM rbs.customers c WHERE c.customerId = ?", customerId)
   if err != nil {
      log.Println("Error fetching addition")
      log.Println(err)
      return
   }
   defer rows.Close()

   for rows.Next() {
      var customer Customer
      rows.Scan(&customer.NAME, &customer.STATE, &customer.PAN, &customer.GUID)
      json.NewEncoder(w).Encode(customer)
      return
   }
}

func main() {
   r := mux.NewRouter()
   r.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
   log.Fatal(http.ListenAndServe(":8000", r))
}
