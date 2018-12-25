package main

import (
    "database/sql"
    _ "gopkg.in/goracle.v2"
    "log"
)

type Customer struct {
   LAST_NAME  string `json:"lastname"`
	 NAME       string `json:"name"`
	 STATE      string `json:"state"`
	 PAN        string `json:"pan"`
   GUID       string `json:"guid"`
}

func getCustomerFromDb(customerId string) Customer {
   var customer Customer
   var conf = getConfig()
   db, err := sql.Open("goracle", conf.DbUsername + "/" + conf.DbPassword + conf.DbName)
   if err != nil {
     log.Println(err)
     return customer
   }
   defer db.Close()

   query := "SELECT last_name, name, state, pan, guid FROM customer WHERE CUSTOMER_ID = " + customerId
   row := db.QueryRow(query)
   err = row.Scan(&customer.LAST_NAME,
                  &customer.NAME,
                  &customer.STATE,
                  &customer.PAN,
                  &customer.GUID)
   if err != nil {
      log.Println(err)
      return customer
   }

   return customer
}
