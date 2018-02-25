package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"./GoCloud"
)

var server GoCloud.Server
var stack []GoCloud.Handle
var port int
var config GoCloud.Config

func main() {
	arg := os.Args[1:]
	if len(arg) > 0 {
		if arg[0] == "server" {
			if len(arg) > 1 {
				if arg[1] == "np" {
					runServerNP()
				} else {
					port, _ = strconv.Atoi(arg[1])
					runServer()
				}
			} else {
				port = 3339
				runServer()
			}
		} else if arg[0] == "run" {
			if len(arg) > 1 {
				runFile(arg[1])
			} else {
				fmt.Println("Error: no hay suficientes parametros para arrancar un archivo")
			}
		} else if arg[0] == "init" {
			initself()
		}
	}
}

func runFile(filename string) {
	data()
	req := new(http.Request)
	req.Method = "GET"
	fmt.Println(GoCloud.Read(filename, config, req, ""))
}

func runServer() {
	data()
	server = GoCloud.Server{stack, config}
	fmt.Println("El servidor a sido iniciado correctamente en el puerto: " + strconv.Itoa(port))
	fmt.Println("Presiona cntrl+c para cerrarlo")
	server.Run(port)
}

func runServerNP() {
	data()
	server = GoCloud.Server{stack, config}
	fmt.Println("El servidor a sido iniciado correctamente sin puerto")
	fmt.Println("Presiona cntrl+c para cerrarlo")
	server.RunNP()
}

func data() {
	plan, _ := ioutil.ReadFile("./Data/config.json")
	json.Unmarshal(plan, &config)
	path := config.Path
	config.Error404.Page = path + "\\Pages\\" + config.Error404.Page
	for _, s := range config.Sites {
		add(s.Handler, path+"\\Pages\\"+s.Page)
	}
}

func add(name string, page string) {
	stack = append(stack, GoCloud.Handle{name, page})
}

func initself() {
	os.Mkdir("Pages", os.ModePerm)
	os.Mkdir("Data", os.ModePerm)
	os.Mkdir("Assets", os.ModePerm)
	os.Mkdir("Assets/JS", os.ModePerm)
	os.Mkdir("Assets/CSS", os.ModePerm)
	f, err := os.Create("Data/config.json")
	if err == nil {
		_, err = f.WriteString(defconfig())
		f.Close()
		if err != nil {
			fmt.Println("Hubo un error escribiendo en config.json")
		}
	} else {
		fmt.Println("Hubo un error creando config.json")
	}
}

func defconfig() string {
	return "{\n" +
		"\"Project\":\"Proyect Name\",\n" +
		"\"Path\":\"\",\n" +
		"\"AllowPHP\":false,\n" +
		"\"PHPPath\":\"\",\n" +
		"\"Saveruntimes\":false,\n" +
		"\"Database\":{\n" +
		"	\"Name\":\"\",\n" +
		"	\"User\":\"\",\n" +
		"	\"Password\":\"\"\n" +
		"},\n" +
		"\"Error404\":{\"Page\":\"The view for a 404 not found error Example:404.html\"},\n" +
		"\"Sites\":[\n" +
		"	{\"Handler\":\"The controller of the view(don't write the '/')\",\"Page\":\"The view Example:index.html\",\"Vars\":[The vars is waiting for],\"Methods\":[The methods Example:GET,POST],\"AlertMethod\":Is a boolean, that if it true the method is use in a variable that will show in your compiled page like 'Method := GET or POST'}\n" +
		"]\n" +
		"}\n"
}
