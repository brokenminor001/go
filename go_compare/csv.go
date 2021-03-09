package main
import (
"encoding/csv"
"fmt"
"os"
)

func main() {

var x []string

  file, err := os.Open("/home/asmolin/Downloads/printers.csv")

  if err != nil {
    panic(err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Comma = ';'

for  {
  record, e := reader.Read()
  if e != nil {
    fmt.Println(e)
    break
  }


  fmt.Println(record[2])

string  = append(x, record[2])

 }


return string 




}
