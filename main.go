package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"

	"log"

	"github.com/apex/gateway"
)

var (
	port = flag.Int("port", -1, "specify a port")
)

func main() {
	flag.Parse()

	// http.Handle("/static/", http.FileServer(http.Dir("static/")))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contact", contactPage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/product", productPage)
	http.HandleFunc("/product-details", productDetailsPage)
	http.HandleFunc("/gallery", galleryPage)

	listenger := gateway.ListenAndServe
	portStr := "n/a"

	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listenger = http.ListenAndServe
	}

	log.Fatal(listenger(portStr, nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func productPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "product.html")
}

func productDetailsPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "product-details.html")
}

func galleryPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "gallery.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
