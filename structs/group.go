package structs

type Group struct {
	Id   float64 `json:"id"`
	Name string  `json:"name"`
}

type GroupResponseList struct {
	Count int     `json:"count"`
	Items []Group `json:"items"`
}

type GroupResponse struct {
	Response GroupResponseList `json:"response"`
}
