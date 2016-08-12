package structs

type Friend struct {
	Id         float64 `json:"id"`
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
}

type FriendResponseList struct {
	Count int      `json:"count"`
	Items []Friend `json:"items"`
}

type FriendResponse struct {
	Response FriendResponseList `json:"response"`
}
