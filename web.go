package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"code.google.com/p/rsc/qr"
	"net/http"
	"strings"
)

var chttp = http.NewServeMux()

type Page struct {
	Title string
	Body string
        Message string
}

func encode(message string) {
	c, err := qr.Encode(message, qr.L)

	if (err != nil) {
		fmt.Println(err)
	} else {
		pngdat := c.PNG()
                s := []string{"images/", message, ".png"}
		ioutil.WriteFile(strings.Join(s,""), pngdat, 0666)
	}

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func homeHandler( w http.ResponseWriter, r *http.Request) {
	if (strings.Contains(r.URL.Path, ".")) {
		chttp.ServeHTTP(w, r)
	} else {
		encodedValue := r.FormValue("encode")
		encode(encodedValue)
		p := &Page{Title: "Ticketbud QR Service", Body: encodedValue, Message: encodedValue}
		renderTemplate(w, "result", p)
	}
}

func main() {
	chttp.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
}
