package controller

import (
  "fmt"
  "github.com/julienschmidt/httprouter"
  "net/http"
)

func StartPage(rw http.ResponseWriter, r *http.Requesst, p httprouter.Params)
{
  text := "Hello"
  fmt.Println(rw. text)
}
