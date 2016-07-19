package main

import "fmt"
import "io"
import "net/http"
import "io/ioutil"

import "strconv"

func (lvcvd *LVCVD) WebDefault(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  url := req.URL
  fmt.Printf("default:\n")
  fmt.Printf("  method: %s\n", req.Method)
  fmt.Printf("  proto:  %s\n", req.Proto)
  fmt.Printf("  scheme: %s\n", url.Scheme)
  fmt.Printf("  host:   %s\n", url.Host)
  fmt.Printf("  path:   %s\n", url.Path)
  fmt.Printf("  frag:   %s\n", url.Fragment)
  fmt.Printf("  body:   %s\n\n", body)

  io.WriteString(w, `{"value":"ok"}`)
}

func (lvcvd *LVCVD) WebAbout(w http.ResponseWriter, req *http.Request) {
  str,e := ioutil.ReadFile("html/about.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (lvcvd *LVCVD) WebInteractive(w http.ResponseWriter, req *http.Request) {
  str,e := ioutil.ReadFile("html/index.html")
  if e!=nil { io.WriteString(w, "error") ; return }
  io.WriteString(w, string(str))
}

func (lvcvd *LVCVD) WebExec(w http.ResponseWriter, req *http.Request) {
  body,err := ioutil.ReadAll(req.Body)
  if err != nil { io.WriteString(w, `{"value":"error"}`); return }

  fmt.Printf("webexec got>>>\n%s\n\n", body)

  rstr,e := lvcvd.JSVMRun(string(body))
  if e!=nil {
    rerr := strconv.Quote(fmt.Sprintf("%v", e))
    io.WriteString(w, `{"value":"error","error":` + rerr + `}`)
    return
  }

  io.WriteString(w, rstr)
}

func (lvcvd *LVCVD) StartSrv() error {
  http.HandleFunc("/", lvcvd.WebDefault)
  http.HandleFunc("/exec", lvcvd.WebExec)
  http.HandleFunc("/about", lvcvd.WebAbout)
  http.HandleFunc("/i", lvcvd.WebInteractive)

  port_str := fmt.Sprintf("%d", lvcvd.Port)

  //e := http.ListenAndServe(":8084", nil)
  e := http.ListenAndServe(":" + port_str, nil)
  return e
}
