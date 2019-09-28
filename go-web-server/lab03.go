package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/RealJK/rss-parser-go"
)

func NewsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()

	var response strings.Builder

	response.WriteString("<div class=\"container\">")
	response.WriteString("<a href=\"http://localhost:3000/static/\">Назад</a>")

	url := r.Form["url"][0]
	rssObject, err := rss.ParseRSS(url)
	if err != nil {

		response.WriteString(fmt.Sprintf("<h1>%s</h1>", rssObject.Channel.Title))

		response.WriteString(fmt.Sprintf("<p>Кол-во новостей: %d</p>", len(rssObject.Channel.Items)))

		for v, item := range rssObject.Channel.Items {
			response.WriteString(fmt.Sprintf("<div>%d.", v+1))

			response.WriteString(fmt.Sprintf("<p>%s</p>", item.Title))
			pattern := regexp.MustCompile(`<!\[CDATA\[\n?.+\n?]]>`)
			response.WriteString(fmt.Sprintf("<p>%s</p>", pattern.ReplaceAllString(item.Description, "$1")))
			response.WriteString("</div>")
		}
	}
	response.WriteString("</div>")
	fmt.Fprintf(w, response.String())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	r.ParseForm()
	var response strings.Builder
	var rssObjects [2]*rss.RSS
	var sites [2]string
	sites[0] = "https://news.mail.ru/rss/90/"
	sites[1] = "https://lenta.ru/rss"
	sum := 0
	response.WriteString("<div class=\"container\">")
	for i := range sites {
		rssObjects[i], _ = rss.ParseRSS(sites[i])
		sum += len(rssObjects[i].Channel.Items)
	}
	m := 0
	arr := make([]rss.Item, sum)
	for i := range sites {
		for _, item := range rssObjects[i].Channel.Items {
			arr[m] = item
			m++
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].PubDate < arr[j].PubDate
	})

	for i := range arr {
		response.WriteString(fmt.Sprintf("<a href=%s><h1>%s\n</h1></a>", arr[i].Link, arr[i].Title))
		response.WriteString(fmt.Sprintf("<h3>%s\n</h3>", arr[i].PubDate))
		response.WriteString(fmt.Sprintf("<h3>%s\n</h3>", arr[i].Author))
		response.WriteString(fmt.Sprintf("<h2>%s\n</h2>", arr[i].Description))
	}
	response.WriteString("</div>")
	fmt.Fprintf(w, response.String())
}

func main() {
	port := ":3000"

	if len(os.Args) > 1 {
		mode := os.Args[1]
		if mode == "-d" || mode == "--mode-dev" {
			println("Running in DEV mode")
			port = "localhost:3000"
		}
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/news", NewsHandler)
	http.HandleFunc("/", HomeHandler)
	fmt.Println("Server listen on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
