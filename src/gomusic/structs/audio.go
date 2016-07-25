package structs


type Audio struct {
	Id int `json:"id"`
	Artist string `json:"artist"`
	Title string `json:"title"`
	Url string `json:"url"`
}

type AudioResponseList struct {
	Count int `json:"count"`
	Items []Audio `json:"items"`
}

type AudioResponse struct {
	Response AudioResponseList `json:"response,omitempty"`
	//Error ErrorResponse `json:"error,omitempty"`
}
