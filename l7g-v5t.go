package main

import "os"
import "fmt"
import "log"
import "io/ioutil"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

import "github.com/abeconnelly/sloppyjson"

type LVCVD struct {
  DB *sql.DB
  HTMLDir string
  JSDir string
  Port int
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

  // add column names to first row
  //
  res_str_array = append(res_str_array, cols)

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
        if local_debug {
          fmt.Printf("raw>>>>\n%v\n", string(raw))
        }

      }
    }

    res_str_array = append(res_str_array, result)

  }

  //DEBUG
  if local_debug {
    fmt.Printf(">>>>\n%v\n", res_str_array)
  }

  return res_str_array, nil
}

func main() {
  local_debug := true
  lvcvd := LVCVD{}

  config_fn := "./l7g-v5t-config.json"
  if len(os.Args)>1 {
    config_fn = os.Args[1]
  }

  config_str,e := ioutil.ReadFile(config_fn)
  if e!=nil { log.Fatal(e) }
  config_json,e := sloppyjson.Loads(string(config_str))
  if e!=nil { log.Fatal(e) }

  lvcvd.Port = int(config_json.O["port"].P)
  lvcvd.HTMLDir = config_json.O["html-dir"].S
  lvcvd.JSDir = config_json.O["js-dir"].S

  err := lvcvd.Init(config_json.O["database"].S)
  if err!=nil { log.Fatal(err) }


  if local_debug {
    fmt.Printf(">> starting\n")
  }

  err = lvcvd.StartSrv()
  if err!=nil { log.Fatal(err) }
}
