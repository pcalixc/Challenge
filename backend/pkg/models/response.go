package models

type QuerySearchResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Person    string `json:"person"`
				Tag       string `json:"email_tag"`
				MessageID string `json:"message_id"`
				From      string `json:"from"`
				To        string `json:"to"`
				Subject   string `json:"subject"`
				Date      string `json:"date"`
				Content   string `json:"content"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
