package structs

type Album struct {
	Id    float64 `json:"id"`
	Title string  `json:"title"`
}

type AlbumResponseList struct {
	Count int     `json:"count"`
	Items []Album `json:"items"`
}

type AlbumResponse struct {
	Response AlbumResponseList `json:"response"`
}
