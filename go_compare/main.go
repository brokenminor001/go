package main
import (
  "github.com/julienschmidt/httprouter"
  "controller"
  "log"
  "net/http"
)


func main() {


  r:=httprouter.New()
  routes(r)


err:=http.ListenAndServe("localhost:4444",r)
if err !=nil{
  log.Fatal(err)
}
}


func routes(r *httprouter.Router) {
  r.ServeFiles("/public/*filepath", http.Dir("public"))
  r.GET("/", controller.StartPage)
  r.GET("/users", controller.GetUsers)
}
