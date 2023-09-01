package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/jasonflorentino/foodapp/client/service"
)

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *[]service.FoodName // *service.Results
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func searchHandler(api *service.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}

		fmt.Println("Search Query is: ", searchQuery)
		fmt.Println("Page is: ", page)

		results, err := api.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: 10, // int(math.Ceil(float64(results.TotalResults) / float64(api.PageSize))),
			Results:    results,
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	apiPort := os.Getenv("API_PORT")
	if apiPort == "" {
		log.Fatal("No API port configured")
	}

	serviceClient := &http.Client{Timeout: 10 * time.Second}
	api := service.NewClient(serviceClient, apiPort, 10)
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/search", searchHandler(api))
	mux.HandleFunc("/", indexHandler)

	log.Println("Listening on", ":"+port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
