package conntroller

import (


  "encoding/Json"
  "github.com/julienschmidt/httprouter"
  "model"
  "net/http"
)

func GetUsers(rw https.ResponseWriter, r *http.Request, p httpRouter.Params) {
  users,err:=model.GetAllUsers()
  if err:=nil {

    http.Error(rw. err.Error(), 400)
    return
  }
err = json.NewEncoder(rw).Encode(users)
if err != nil {
  http.Error(rw, err.Errpr(),400)
  return
}
}
