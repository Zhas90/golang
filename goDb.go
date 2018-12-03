package main

import (
    "database/sql"
    _ "gopkg.in/goracle.v2"
    "github.com/gorilla/mux"
    "encoding/json"
    "net/http"
    "log"
    "time"
)

type Customer struct {
  ID                     string `json:"id"`
	PHONE_NUMBER           string `json:"phoneNumber"`
  NAME                   string `json:"name"`
  IIN                    string `json:"iin"`
  STATE                  string `json:"state"`
  PAN                    string `json:"pan"`
  GUID                   string `json:"guid"`
  LAST_NAME              string `json:"lastName"`
  DOB                    time.Time   `json:"dob"`
  CUSTOMER_ID            string `json:"customerId"`
  SEC_ANSWER             string `json:"secAnswer"`
  SEC_QUESTION_ID        string `json:"secQuestionId"`
  LANG_DEF               string `json:"langDef"`
  PIN_UPDATE             time.Time   `json:"pinUpdate"`
  SECOND_FACTOR_REQ      string `json:"secFactorReq"`
  TYPE_OF_SECOND_FACTOR  string `json:"typeOfSecondFactor"`
  TERMS                  string `json:"terms"`
  EMAIL                  string `json:"email"`
  LAST_TRUE_ANS_DATE     time.Time   `json:"lastTrueAnsDate"`
  EMAIL_NOTIFICATION     string `json:"emailNotification"`
  SMS_NOTIFICATION       string `json:"smsNotification"`
  MIDDLE_NAME            string `json:"middleName"`
  FIRST_TRANSACTION      string `json:"firstTransaction"`
  FIRST_SESSION          string `json:"firstSession"`
  ADDRESS                string `json:"address"`
  BRANCH                 string `json:"branch"`
}

func main() {
   r := mux.NewRouter()
   r.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
   log.Fatal(http.ListenAndServe(":8000", r))
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json")

   db, err := sql.Open("goracle", "RBS/p9NFDS4qFdU8@172.24.10.133:1521/wmr")
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
