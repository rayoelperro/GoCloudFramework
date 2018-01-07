package main
import (
"fmt"
"./GoCloud"
)
func main(){
POST_nombre := "Arnaldo"
POST_apellido := "Elorza"
POST_edad := "45"
Method := "POST"
fmt.Println("<!DOCTYPE html>")
fmt.Println("<html lang=\"en\">")
fmt.Println("<head>")
fmt.Println("<meta charset=\"UTF-8\">")
fmt.Println("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
fmt.Println("<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">")
fmt.Println("<title>Your new user already create</title>")
fmt.Println("</head>")
fmt.Println("<body>")
if Method == "POST"{
table := GetDBTable("users")
table.Insert([]string{POST_nombre,POST_apellido,POST_edad})
fmt.Println(" <h1>"+POST_nombre+" "+POST_apellido+" de "+POST_edad+" años se ha insertado correctamente</h1>")
}else{
fmt.Println("<h1>No se han ofrecido parametros</h1>")
}
fmt.Println("</body>")
fmt.Println("</html>")
}
func GetDBTable(tablename string) GoCloud.DataTable {
db := GoCloud.Connect("root","","golang")
return db.GetTable(tablename)
}
