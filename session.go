package main

import (
    "database/sql"
    _ "gopkg.in/goracle.v2"
    "log"
)

type Session struct {
   ID  string `json:"id"`
	 IP  string `json:"ip"`
	 OS  string `json:"os"`
}

func getSessionFromDb(sessionId string) Session {
   var session Session
   var conf = getConfig()
   db, err := sql.Open("goracle", conf.DbUsername + "/" + conf.DbPassword + conf.DbName)
   if err != nil {
     log.Println(err)
     return session
   }
   defer db.Close()

   query := "SELECT id, ip, os FROM sessions_ WHERE SESSION_ID = " + sessionId
   row := db.QueryRow(query)
   err = row.Scan(&session.ID,
                  &session.IP,
                  &session.OS)

   if err != nil {
      log.Println(err)
      return session
   }

   return session
}
