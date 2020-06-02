package parsing

import (
	"homelibrary/config"
	"strings"
)

// Converts url path to directory path
func MediaURL(url string) string {
	if (config.MediaRoot[len(config.MediaRoot)-1:] == "/" && url[0:1] == "/") {
		url = url[1:len(url)]
	}
	return config.MediaRoot+url
}

// Checks if the end of Name is .pdf
func NameIsPDF(Name string) bool {
	return strings.Contains(Name,".pdf")
	if len(Name) <= 4 {
		return false
	}
	return Name[len(Name)-4:] == ".pdf"
}

