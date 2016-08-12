package structs

import "net/url"

type Audio struct {
	Id     float64 `json:"id"`
	Artist string  `json:"artist"`
	Title  string  `json:"title"`
	Url    string  `json:"url"`
}

func (a *Audio) CleanUrl() string {
	parsed, _ := url.Parse(a.Url)
	return parsed.Scheme + "://" + parsed.Host + parsed.Path
}

type AudioResponseList struct {
	Count int     `json:"count"`
	Items []Audio `json:"items"`
}

type AudioResponse struct {
	Response AudioResponseList `json:"response"`
}
