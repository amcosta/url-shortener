/*
Sistema encurtador de url

faça uma requisição http e recebe uma url encurtada, estamos trabalhando para gerar métricas de acesso da sua url

Work In Progress
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"amcosta.dev/encurtador-url/repository"
	"amcosta.dev/encurtador-url/services"
	"amcosta.dev/encurtador-url/shortlink"
)

// Comentando a estrutura
type UrlPostRequest struct {
	Url string
}

type Url struct {
	Id           string `json:"id"`
	OriginalLink string `json:"original_link"`
	ShortedLink  string `json:"shorted_link"`
}

func main() {
	http.HandleFunc("/urls", createUrl)
	http.ListenAndServe(":8080", nil)
}

func createUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Metodo não permitido")
		return
	}

	var link UrlPostRequest
	json.NewDecoder(r.Body).Decode(&link)

	urlRepository := &repository.InMemory{}
	urlService := services.NewUrlService(urlRepository, shortlink.Generator{})

	url := urlService.Create(link.Url)
	jsonData, err := json.Marshal(url)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(jsonData))
}
