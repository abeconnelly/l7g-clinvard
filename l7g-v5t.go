package main

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

type LVCVD struct {
  DB *sql.DB
}


func (lvcvd *LVCVD) Init(sql_fn string) error {
  var err error
  lvcvd.DB, err = sql.Open("sqlite3", sql_fn)
  if err !=nil { panic(err) }
  return nil
}

func (lvcvd *LVCVD) SQLExec(req string) ([][]string, error ) {
  local_debug := true

  if local_debug {
    fmt.Printf("SQLExec: request: \"%s\"\n", req)
  }

  rows,err := lvcvd.DB.Query(req)
  if err!=nil { return nil, err }
  cols,e := rows.Columns() ; _ = cols
  if e!=nil { return nil, e }

  rawResult := make([][]byte, len(cols))

  res_str_array := [][]string{}

  dest := make([]interface{}, len(cols))
  for i,_ := range rawResult {
    dest[i] = &rawResult[i]
  }

  for rows.Next() {
    err := rows.Scan(dest...)
    if err!=nil { return nil,err }

    result := make([]string, len(cols))

    for i,raw := range rawResult {
      if raw==nil {
        result[i] = "\n"
      } else {
        result[i] = string(raw)

        //DEBUG
        fmt.Printf("raw>>>>\n%v\n", string(raw))

      }
    }

    res_str_array = append(res_str_array, result)

  }

  //DEBUG
  fmt.Printf(">>>>\n%v\n", res_str_array)

  return res_str_array, nil
}

func main() {
  local_debug := true
  lvcvd := LVCVD{}

  err := lvcvd.Init("./l7g-v5t-clinvar.sqlite3")
  if err!=nil { panic(err) }

  if local_debug {
    fmt.Printf(">> starting\n")
  }

  err = lvcvd.StartSrv()
  if err!=nil { panic(err) }

}
