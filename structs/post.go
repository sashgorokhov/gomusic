package structs

type Attachment struct {
	Type  string `json:"type"`
	Audio Audio  `json:"audio"`
}

type Post struct {
	Id           float64      `json:"id"`
	Owner_id     float64      `json:"owner_id"`
	Post_type    string       `json:"post_type"`
	Text         string       `json:"text"`
	Attachments  []Attachment `json:"attachments"`
	Copy_history []Post       `json:"copy_history"`
}

type PostResponseList struct {
	Count int    `json:"count"`
	Items []Post `json:"items"`
}

type PostResponse struct {
	Response PostResponseList `json:"response"`
}
