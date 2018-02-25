package GoCloud

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	Project      string
	Path         string
	AllowPHP     bool
	PHPPath      string
	Saveruntimes bool
	Database     Database
	Error404     Site
	Sites        []Site
}

type Database struct {
	Name     string
	User     string
	Password string
}

type Site struct {
	Handler     string
	Page        string
	Vars        []string
	Methods     []string
	AlertMethod bool
}

type Handle struct {
	Name string
	Page string
}

type Server struct {
	Handlers []Handle
	Props    Config
}

func (s Server) Run(port int) {
	apileHandlers(s)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func (s Server) RunNP() {
	apileHandlers(s)
	http.ListenAndServe("", nil)
}

func apileHandlers(s Server) {
	http.Handle("/Assets/", http.StripPrefix("/Assets/", http.FileServer(http.Dir("Assets"))))
	for _, hand := range s.Handlers {
		name := hand.Name
		page := hand.Page
		http.HandleFunc("/"+name, func(r http.ResponseWriter, a *http.Request) {
			if a.URL.Path != "/"+name {
				errorHandler(r, a, http.StatusNotFound, "/"+name, s)
				return
			} else {
				t := time.Now().UTC().String()
				fmt.Print(t + " - ")
				fmt.Println("Llamando al archivo en: " + page)
				io.WriteString(r, Read(page, s.Props, a, name))
			}
		})
	}
}

func (g *Server) Handler(name string, page string) {
	g.Handlers = append(g.Handlers, Handle{name, page})
}

func errorHandler(r http.ResponseWriter, a *http.Request, status int, notfound string, s Server) {
	r.WriteHeader(status)
	if status == http.StatusNotFound {
		t := time.Now().UTC().String()
		fmt.Print(t + " - ")
		fmt.Println("Llamando al archivo de error 404 al intentar cargar: " + notfound)
		io.WriteString(r, Read(s.Props.Error404.Page, s.Props, a, s.Props.Error404.Handler))
	}
}
