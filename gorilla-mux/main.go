package main

import (
  _"fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/api/say/hello", homeHandler)
  router.PathPrefix("/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
  http.ListenAndServe(":8100", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("අායුබෝවන්"))
}
