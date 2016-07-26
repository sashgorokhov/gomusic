package formatters

import (
	"gomusic/structs"
	"strings"
	"strconv"
)


var Album_format_choices = []string{"id", "title"}
const Album_format_default = "id,title"

func format_album_title(title string, replace_chars bool) string {
	return title
}

func Format_album(album *structs.Album, format_string string, quiet bool, replace_chars bool) string {
	if quiet {
		return strconv.Itoa(album.Id)
	}
	if format_string == "" {
		format_string = Album_format_default
	}
	var colums []string
	for _, v := range strings.Split(format_string, ",") {
		switch  {
			case v == "id": {
				colums = append(colums, strconv.Itoa(album.Id))
			}
			case v == "title": {
				colums = append(colums, format_album_title(album.Title, replace_chars))
			}
		}
	}
	return strings.Join(colums, " ")
}
