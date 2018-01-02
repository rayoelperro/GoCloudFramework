package main
import (
"fmt"
"./GoCloud"
)
func main(){

fmt.Println("<!DOCTYPE html>")
fmt.Println("<html lang=\"en\">")
fmt.Println("<head>")
fmt.Println("<meta charset=\"UTF-8\">")
fmt.Println("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
fmt.Println("<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">")
fmt.Println("<title>Mi Pagina web</title>")
fmt.Println("</head>")
fmt.Println("<body>")
table := GetDBTable("users")
var res [][]string
res = table.QueryAll([]string{"nombre","apellido","edad"})
for _,put := range res{
fmt.Println("<h1>Hola "+put[0]+" "+put[1]+" tienes "+put[2]+" años</h1>")
}
fmt.Println("</body>")
fmt.Println("</html>")
}
func GetDBTable(tablename string) GoCloud.DataTable {
db := GoCloud.Connect("root","","golang")
return db.GetTable(tablename)
}
