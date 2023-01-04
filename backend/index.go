package main

import (
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

func list_all(folder_name string) []string {
	var folder_list []string

	outputDirRead, _ := os.Open(folder_name)

	outputDirFiles, _ := outputDirRead.Readdir(0)

	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		outputNameHere := outputFileHere.Name()
		folder_list = append(folder_list, outputNameHere)
	}
	return folder_list
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

func MailContent(path string) string {
	sys_file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	text := string(sys_file)

	return text
}

func KK(path string) string {
	var mail string

	if IsPath(path) {
		folders2 := list_all(path)
		for f2 := range folders2 {
			KK(path + "/" + folders2[f2])
		}
	} else {
		fmt.Println("Indexing:   ", path)
		mail = MailContent(path)

		parse_data(mail)

		IndexData(*parse_data(mail))

	}
	return mail
}

type Email struct {
	MessageId string `json:"message_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Date      string `json:"date"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

func parse_data(content string) *Email {
	email := &Email{}
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

func IndexData(data Email) {

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
}

func main() {
	user_list := list_all("../enron_mail_20110402/maildir")

	for u := range user_list {
		folders := list_all("../enron_mail_20110402/maildir/" + user_list[u])

		for f := range folders {
			KK("../enron_mail_20110402/maildir/" + user_list[u] + "/" + folders[f])

		}
	}
}
