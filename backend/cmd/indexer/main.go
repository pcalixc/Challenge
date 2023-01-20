package main

import (
	"backend/pkg/services"
	"log"
)

func main() {
	path := "../../../enron_mail_20110402/maildir/"

	user_list, err := services.ListFiles(path)
	if err != nil {
		log.Printf("Error while indexing email: %v", err)
		return
	}

	for u := range user_list {
		folders, err := services.ListFiles(path + user_list[u])
		if err != nil {
			log.Printf("Error while listing email: %v", err)
			return
		}

		for f := range folders {
			services.IndexEmail(path + user_list[u] + "/" + folders[f])

		}
	}
}
