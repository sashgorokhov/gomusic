package formatters

import (
	"github.com/sashgorokhov/gomusic/structs"
	"strconv"
	"strings"
)

var Audio_format_choices = []string{"id", "title", "url"}
const Audio_format_default = "id,url,title"

func format_audio_title(artist, title string) string {
	return strings.TrimSpace(artist) + " - " + strings.TrimSpace(title)
}

func Format_audio_filename(audio *structs.Audio) string {
	return format_audio_title(audio.Artist, audio.Title) + ".mp3"
}

func Format_audio(audio *structs.Audio, format_string string) string {
	if format_string == "" {
		format_string = Audio_format_default
	}
	var colums []string
	for _, v := range strings.Split(format_string, ",") {
		switch  {
			case v == "id": {
				colums = append(colums, strconv.Itoa(audio.Id))
			}
			case v == "title": {
				colums = append(colums, format_audio_title(audio.Artist, audio.Title))
			}
			case v == "url": {
				colums = append(colums, audio.CleanUrl())
			}
		}
	}
	return strings.Join(colums, " ")
}