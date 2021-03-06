package formatters

import (
	"github.com/sashgorokhov/gomusic/structs"
	"strconv"
	"strings"
)

var Friend_format_choices = []string{"id", "name"}

const Friend_format_default = "id,name"

func format_friend_name(first_name, last_name string) string {
	return strings.TrimSpace(first_name) + " " + strings.TrimSpace(last_name)
}

func Format_friend(friend *structs.Friend, format_string string) string {
	if format_string == "" {
		format_string = Friend_format_default
	}
	var colums []string
	for _, v := range strings.Split(format_string, ",") {
		switch {
		case v == "id":
			{
				colums = append(colums, strconv.Itoa(int(friend.Id)))
			}
		case v == "name":
			{
				colums = append(colums, format_friend_name(friend.First_name, friend.Last_name))
			}
		}
	}
	return strings.Join(colums, " ")
}
