package models

type QuerySearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Email []struct {
			Index     string `json:"_index"`
			ID        string `json:"_id"`
			Timestamp string `json:"@timestamp"`
			Source    struct {
				MessageID string `json:"message_id"`
				From      string `json:"from"`
				To        string `json:"to"`
				Subject   string `json:"subject"`
				Date      string `json:"date"`
				Content   string `json:"content"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}
