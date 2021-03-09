package main
import (

"fmt"
 "database/sql"
 _ "github.com/lib/pq"
)
type groupe struct {
id int
name string
points int
}
type team struct {

  name string
}



var tone string


func main() {
connStr := "user=anb password=oadnpvia host=130.61.60.226 dbname=anb sslmode=disable"
db, err := sql.Open("postgres", connStr)
if err != nil {
  panic(err)
}

defer db.Close()

rows, err := db.Query("select name, points from groupe_a")
if err !=nil {
  panic(err)
}
defer rows.Close()
groupes := []groupe{}

for rows.Next() {
  g := groupe{}
  err := rows.Scan(&g.name, &g.points)
  if err != nil {
    fmt.Println(err)
    continue
  }
  groupes = append(groupes, g)
  }
for _, g := range groupes{
  fmt.Println(g.name[0])


}




}
