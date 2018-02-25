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
fmt.Println("<link rel=\"stylesheet\" type=\"text/css\" href=\"/Assets/CSS/design.css\">")
fmt.Println("<script src=\"/Assets/JS/main.js\"></script>")
fmt.Println("<title>Mi Pagina web</title>")
fmt.Println("</head>")
fmt.Println("<body>")
fmt.Println("<div class=\"main\">")
fmt.Println("Holaaaaaa!")
fmt.Println("<h1 class=\"title\">Escribe un nombre para saludar:</h1>")
fmt.Println("<form action=\"/greet\" method=\"post\" class=\"topper\">")
fmt.Println("<input type=\"text\" name=\"nombre\" id=\"nombre\">")
fmt.Println("<input type=\"submit\" value=\"Aceptar\">")
fmt.Println("</form>")
fmt.Println("</br>")
fmt.Println("<a href=\"/newuser\">Registrar nuevo usuario</a>")
fmt.Println("</br>")
fmt.Println("<img src=/Assets/IMG/welcome.png></img>")
fmt.Println("</div>")
fmt.Println("</body>")
fmt.Println("</html>")
}
func GetDBTable(tablename string) GoCloud.DataTable {
db := GoCloud.Connect("root","","golang")
return db.GetTable(tablename)
}
