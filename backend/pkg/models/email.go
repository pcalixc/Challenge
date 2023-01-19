package models

type Email struct {
	MessageId string `json:"message_id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Date      string `json:"date"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}
