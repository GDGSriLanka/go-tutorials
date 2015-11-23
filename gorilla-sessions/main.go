package main

import (
  _"fmt"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
  "io"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8100", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  session, err := store.Get(r,"session-name")
  if err != nil {
    http.Error(w, err.Error(),500)
    return
  }
  session.Values["foo"] = "bar"
  session.Values[42] = 43
  session.Save(r,w)
  str, _ := session.Values["foo"].(string)
  io.WriteString(w,str)
}
