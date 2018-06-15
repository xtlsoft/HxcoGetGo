package main

import (
	"github.com/xtlsoft/router"
	"flag"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func getContent(name string) string {

	f, _ := os.Open(name)
	defer f.Close()
	r, _ := ioutil.ReadAll(f)

	return string(r)

}

func main() {

	port := flag.String("p", ":21380", "Port to listen.")
	flag.Parse()

	r := router.New()
	r.Group("", func (g *router.Group){
		g.Get("/{@ident:software}", func (req router.Request) router.Response {
			s := req.(*router.HttpRequest).RouterVariable["software"]
			soft := Softwares[s]
			return &router.HttpResponse{
				StatusCode: 200,
				Body: fmt.Sprintf(getContent("./template/software.html"), 
					soft.Name, 
					soft.Name, 
					soft.Windows64, 
					soft.Windows32, 
					soft.MacOsX, 
					soft.IPhone, 
					soft.IPad, 
					soft.Android),
			}
		})
		g.Get("/", func (req router.Request) router.Response {
			return &router.HttpResponse{
				StatusCode: 200,
				Body: getContent("./template/index.html"),
			}
		})
		g.Get("/index.html", func (req router.Request) router.Response {
			return &router.HttpResponse{
				StatusCode: 200,
				Body: getContent("./template/index.html"),
			}
		})
		g.Get("/mobile.html", func (req router.Request) router.Response {
			return &router.HttpResponse{
				StatusCode: 200,
				Body: getContent("./template/mobile.html"),
			}
		})
	})

	http.Handle("/elements/", http.StripPrefix("/elements/", http.FileServer(http.Dir("./assets"))))

	router.HttpServe(r, *port)

}