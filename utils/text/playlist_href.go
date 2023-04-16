package text

import (
	"strings"
)

func GetPlaylistHref(link string) string {
	s := strings.Split(link, "/")
	return "spotify:playlist:" + s[5]
}
