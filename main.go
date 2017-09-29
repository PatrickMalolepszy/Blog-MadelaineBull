package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var t *template.Template

type header struct {
	Title    string
	SubTitle string
	Picture  string
}

type layout struct {
	Header header
}

func init() {
	t = template.Must(template.ParseGlob("templates/*"))
	fmt.Println("templates loaded")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t.ParseFiles("content/home.html")
	indexLayout := layout{
		Header: header{
			Title:    "Madelaine Bull",
			SubTitle: "Personal Website",
			Picture:  "home-bg.jpg",
		},
	}
	err := t.ExecuteTemplate(w, "layout", indexLayout)
	if err != nil {
		fmt.Println(err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t.ParseFiles("content/about.html")
	aboutLayout := layout{
		Header: header{
			Title:    "About Me",
			SubTitle: "",
			Picture:  "",
		},
	}
	err := t.ExecuteTemplate(w, "layout", aboutLayout)
	if err != nil {
		fmt.Println(err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	t.ParseFiles("content/researchPaper.html")
	postLayout := layout{
		Header: header{
			Title:    "",
			SubTitle: "Understanding Native Ecocriticism in the Context of Canadian Identity and Environmental Threats",
			Picture:  "",
		},
	}
	err := t.ExecuteTemplate(w, "layout", postLayout)
	if err != nil {
		fmt.Println(err)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	t.ParseFiles("content/contact.html")
	postLayout := layout{
		Header: header{
			Title:    "Contact Me",
			SubTitle: "",
			Picture:  "",
		},
	}
	err := t.ExecuteTemplate(w, "layout", postLayout)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	fmt.Printf("starting server\n")

	// routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about.html", aboutHandler)
	http.HandleFunc("/researchPaper.html", postHandler)
	http.HandleFunc("/contact.html", contactHandler)

	// load assets
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("vendor"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
