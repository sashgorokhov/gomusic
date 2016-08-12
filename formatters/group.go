package formatters

import (
	"github.com/sashgorokhov/gomusic/structs"
	"strconv"
	"strings"
)

var Group_format_choices = []string{"id", "name"}

const Group_format_default = "id,name"

func format_group_name(name string) string {
	return strings.TrimSpace(name)
}

func Format_group(group *structs.Group, format_string string) string {
	if format_string == "" {
		format_string = Group_format_default
	}
	var colums []string
	for _, v := range strings.Split(format_string, ",") {
		switch {
		case v == "id":
			{
				colums = append(colums, strconv.Itoa(int(group.Id)))
			}
		case v == "name":
			{
				colums = append(colums, format_group_name(group.Name))
			}
		}
	}
	return strings.Join(colums, " ")
}
