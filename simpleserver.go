package main

import (
  "net/http"
)

func main() {

  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8100", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("අායුබෝවන්"))
}
