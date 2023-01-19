package main

import (
	"backend/pkg/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func ListFiles(folderName string) ([]string, error) {
	var folderList []string

	outputDirRead, err := os.Open(folderName)
	if err != nil {
		return nil, fmt.Errorf("could not open folder %s: %v", folderName, err)
	}
	defer outputDirRead.Close()

	outputDirFiles, err := outputDirRead.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("could not read files from folder %s: %v", folderName, err)
	}

	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		outputNameHere := outputFileHere.Name()
		folderList = append(folderList, outputNameHere)
	}

	return folderList, nil
}

func IsPath(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

func GetMailContent(path string) string {
	sys_file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(sys_file)
	return content
}

func IndexEmail(path string) (string, error) {
	var mail string

	if IsPath(path) {
		folders2, err := ListFiles(path)
		if err != nil {
			return "", fmt.Errorf("error while listing files in path %s: %v", path, err)
		}
		for f2 := range folders2 {
			_, err := IndexEmail(path + "/" + folders2[f2])
			if err != nil {
				return "", fmt.Errorf("error while indexing email in path %s: %v", path, err)
			}
		}
	} else {
		fmt.Println("Indexing:   ", path)
		mail = GetMailContent(path)

		IndexData(*ParseData(mail))
	}
	return mail, nil
}

func ParseData(content string) *models.Email {
	email := &models.Email{}
	contentFile := strings.SplitN(content, "\r\n\r\n", 2)
	emailDetails := strings.Split(contentFile[0], "\r\n")

	for _, emailDetail := range emailDetails {
		detail := strings.SplitN(emailDetail, ":", 2)

		switch detail[0] {
		case "Message-ID":
			email.MessageId = detail[1]
		case "From":
			email.From = detail[1]
		case "To":
			email.To = detail[1]
		case "Date":
			email.Date = detail[1]
		case "Subject":
			email.Subject = detail[1]
		default:
			continue
		}
	}
	email.Content = contentFile[1]

	return email
}

func IndexData(data models.Email) error {

	newData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(newData)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/Enron1/_doc", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "t5w0109%")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Body: " + string(body))

	return nil
}

func main() {
	path := "../enron_mail_20110402/maildir/"

	user_list, err := ListFiles(path)
	if err != nil {
		log.Printf("Error while indexing email: %v", err)
		return
	}

	for u := range user_list {
		folders, err := ListFiles(path + user_list[u])
		if err != nil {
			log.Printf("Error while listing email: %v", err)
			return
		}

		for f := range folders {
			IndexEmail(path + user_list[u] + "/" + folders[f])

		}
	}
}
