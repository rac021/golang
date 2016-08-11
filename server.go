package main

import (
  "encoding/json"
  "net/http"
  "fmt"
)

type Data struct {
    Name  string
    Age   int
}

func rootHandler(writer http.ResponseWriter, request *http.Request )  {
  
  //writer.Write([]byte("\n"))
  //writer.Write([]byte("Hello from Go server "))
  //writer.Write([]byte("\n"))
  
  data := Data{ Name: "toto", Age : 21 }
  payload, _ := json.Marshal(data)

  writer.Header().Add("Content-Type", "application/json" )

  //writer.Write([]byte("\n"))
  writer.Write(payload)
  writer.Write([]byte("\n"))
  writer.Write([]byte("\n"))
}

func main() {  
  fmt.Printf("Server test")
  http.HandleFunc("/", rootHandler)
  http.ListenAndServe(":8080", nil )  
}

