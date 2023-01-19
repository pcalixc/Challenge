package handlers

import (
	"backend/pkg/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	URL_BASE_ZINC = "http://localhost:4080/api"
	INDEX         = "Enron1"
)

func SearchRecord(term string) (*models.QuerySearchResponse, error) {
	client := &http.Client{}
	queryResponse := &models.QuerySearchResponse{}
	typeSearch := "matchphrase"
	if term == "" {
		typeSearch = "matchall"
	}
	queryRequest := &models.QuerySearchRequest{
		SearchType: typeSearch,
		Query: models.QueryTerm{
			Term: term,
		},
		From:      0,
		MaxResult: 5,
		Source:    []string{},
	}
	query, _ := json.Marshal(queryRequest)
	request, err := http.NewRequest("POST", URL_BASE_ZINC+"/"+INDEX+"/_search", strings.NewReader(string(query)))
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth("admin", "t5w0109%")
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &queryResponse); err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, errors.New("Error")
	}
	return queryResponse, nil
}

func ZincSearch(mux chi.Router) {
	mux.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		phrase := r.URL.Query().Get("term")
		queryResponse, err := SearchRecord(phrase)
		if err != nil {
			render.JSON(w, r, http.StatusInternalServerError)
		}
		render.JSON(w, r, queryResponse)
	})
}
