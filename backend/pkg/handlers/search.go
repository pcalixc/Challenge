package handlers

import (
	"backend/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

const (
	ZincUrl   = "http://localhost:4080/api"
	IndexName = "Enron1"
)

func SearchEmailsWithFilter(term string) (*models.QuerySearchResponse, error) {

	godotenv.Load("../.env")

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
		MaxResult: 3,
		Source:    []string{},
	}
	query, _ := json.Marshal(queryRequest)
	request, err := http.NewRequest("POST", ZincUrl+"/"+IndexName+"/_search", strings.NewReader(string(query)))
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth("admin", os.Getenv("ZINC_PSW"))
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

func SearchEmail(w http.ResponseWriter, r *http.Request) {

	// urlPath := r.URL.Path

	// parts := strings.Split(urlPath, "/")

	// term := parts[len(parts)-1]

	//term := r.URL.Query().Get("term")

	temp := chi.URLParam(r, "temp")

	fmt.Println(r, " valor de r")
	queryResponse, err := SearchEmailsWithFilter(temp)
	if err != nil {
		render.JSON(w, r, http.StatusInternalServerError)
	}

	render.JSON(w, r, queryResponse)

}
