package GoCloud

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Read(page string, props Config, a *http.Request, name string) string {
	f, _ := os.Create("inruntime.go")
	f.WriteString(SumCode(page, props, a, name))
	f.Close()
	p, e := exec.LookPath("go")
	if e == nil {
		var sherr bytes.Buffer
		cmd := exec.Command(p, "run", props.Path+"\\"+"inruntime.go")
		cmd.Stderr = &sherr
		out, err := cmd.Output()
		if props.Saveruntimes {
			if _, err := os.Stat(props.Path + "\\" + "runtimes"); os.IsNotExist(err) {
				os.Mkdir(props.Path+"\\"+"runtimes", os.ModePerm)
			}
			t := time.Now()
			name := strconv.Itoa(t.Hour()) + "_" + strconv.Itoa(t.Minute()) + "_" + strconv.Itoa(t.Second()) + "_" + strconv.Itoa(t.Nanosecond())
			name += "-" + strconv.Itoa(t.Day()) + "_" + t.Month().String() + "_" + strconv.Itoa(t.Year())
			os.Rename(props.Path+"\\"+"inruntime.go", props.Path+"\\"+"runtimes\\"+name)
		}
		if err != nil {
			fmt.Println("Ha habído un error cargando la página: " + err.Error() + " " + sherr.String())
			return "Error cargando la página: " + sherr.String()
		} else {
			return string(out)
		}
	} else {
		return "Error buscando el compilador"
	}
}

func SumCode(page string, props Config, a *http.Request, name string) string {
	f, _ := os.Open(page)
	scanner := bufio.NewScanner(f)
	header := "package main\nimport (\n\"fmt\"\n\"./GoCloud\"\n"
	final := ")\nfunc main(){\n"
	v := Vars(a, props, name)
	final += strings.Join(v, "\n") + "\n"
	for scanner.Scan() {
		text := scanner.Text()
		for len(text) > 0 && string([]rune(text)[0]) == " " || len(text) > 0 && string([]rune(text)[0]) == "	" {
			text = text[1:]
		}
		if len(text) > 0 {
			if string([]rune(text)[0]) == "#" {
				final += text[1:] + "\n"
			} else if string([]rune(text)[0]) == "@" {
				header += "\"" + text[1:] + "\"\n"
			} else if string([]rune(text)[0]) == "&" {
				parts := strings.Split(text[1:], ":")
				if len(parts) == 2 {
					switch parts[0] {
					case "js":
						final += "fmt.Println(\"<script src=\\\"/Assets/JS/" + parts[1] + "\\\"></script>\")\n"
					case "css":
						final += "fmt.Println(\"<link rel=\\\"stylesheet\\\" type=\\\"text/css\\\" href=\\\"/Assets/CSS/" + parts[1] + "\\\">\")\n"
					default:
						final += "panic(\"Tipo de recurso desconocido\")\n"
					}
				} else {
					final += "panic(\"Demasiados argumentos para cargar un recurso\")\n"
				}
			} else if string([]rune(text)[0]) == "?" {
				if !OneStartsWith(v, text[1:]) {
					final += text[1:] + " := \"\"\n"
				}
			} else if string([]rune(text)[0]) == "$" {
				add := text[1:]
				add = strings.Replace(add, "\"", "\\\"", -1)
				add = strings.Replace(add, "'", "\"", -1)
				add = strings.Replace(add, "#(", "\"+", -1)
				add = strings.Replace(add, ")#", "+\"", -1)
				final += "fmt.Println(\"" + add + "\")\n"
			} else {
				add := strings.Replace(text, "\"", "\\\"", -1)
				final += "fmt.Println(\"" + add + "\")\n"
			}
		}
	}
	final += "}\n"
	final = header + final + GetDBTable(props)
	f.Close()
	return final
}

func GetDBTable(a Config) string {
	return "func GetDBTable(tablename string) GoCloud.DataTable {\ndb := GoCloud.Connect(\"" + a.Database.User + "\",\"" + a.Database.Password + "\",\"" + a.Database.Name + "\")\n" +
		"return db.GetTable(tablename)\n}\n"
}

func Vars(a *http.Request, props Config, name string) []string {
	var res []string
	if a.Method == "POST" {
		fmt.Println("El metodo es POST")
		for _, nom := range props.Sites {
			if nom.Handler == name {
				if nom.Vars != nil && Contains(nom.Methods, "POST") {
					for _, nv := range nom.Vars {
						fmt.Println("POST_" + nv + " := \"" + a.PostFormValue(nv) + "\"")
						res = append(res, "POST_"+nv+" := \""+a.PostFormValue(nv)+"\"")
					}
				}
				break
			}
		}
	} else if a.Method == "GET" {
		fmt.Println("El metodo es GET")
		for _, nom := range props.Sites {
			if nom.Handler == name {
				if nom.Vars != nil && Contains(nom.Methods, "GET") {
					for _, nv := range nom.Vars {
						fmt.Println("GET_" + nv + " := \"" + a.FormValue(nv) + "\"")
						res = append(res, "GET_"+nv+" := \""+a.FormValue(nv)+"\"")
					}
				}
				break
			}
		}
	} else {
		fmt.Println("El metodo: " + a.Method + " no esta registrado")
	}
	for _, nom := range props.Sites {
		if nom.Handler == name {
			if nom.AlertMethod {
				fmt.Println("Method := \"" + a.Method + "\"")
				res = append(res, "Method := \""+a.Method+"\"")
			}
			break
		}
	}
	return res
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func OneStartsWith(s []string, e string) bool {
	for _, a := range s {
		if strings.HasPrefix(a, e) {
			return true
		}
	}
	return false
}
