package formatters

import (
	"github.com/sashgorokhov/gomusic/structs"
	"strconv"
	"strings"
)

var Album_format_choices = []string{"id", "title"}

const Album_format_default = "id,title"

func format_album_title(title string) string {
	return title
}

func Format_album(album *structs.Album, format_string string) string {
	if format_string == "" {
		format_string = Album_format_default
	}
	var colums []string
	for _, v := range strings.Split(format_string, ",") {
		switch {
		case v == "id":
			{
				colums = append(colums, strconv.Itoa(int(album.Id)))
			}
		case v == "title":
			{
				colums = append(colums, format_album_title(album.Title))
			}
		}
	}
	return strings.Join(colums, " ")
}
