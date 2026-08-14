package main

import(
	"net/http"
	"html/template"
)


func HomeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil{
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}



func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	if text == "" || bannerName == "" {
		http.Error(w, "Missing input", http.StatusBadRequest)
		return
	}

	banner, err := LoadBanner(bannerName + ".txt")
	if err != nil {
		http.Error(w, "Banner not found", http.StatusNotFound)
		return
	}

	lines := ParseInput(text)

	result := RenderToString(lines, banner)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Result string
	}{
		Result: result,
	}

	err = tmpl.Execute(w, data)
	if err != nil{
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}
}
