package main

import (
    "database/sql"
    "fmt"
    _ "gopkg.in/goracle.v2"
)

func main() {
    dsn := `system/manager@//1.2.3.4:1521/xe`

    _, err := sql.Open("goracle", dsn)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("successfully connected to oracle server using dsn:", dsn)
}
