package helper

import "strings"

func ParseTemplateFilename(requestPath string) string {
	filename := requestPath

	if strings.HasSuffix(filename, "/") {
		filename = filename + "index.html"
	}

	return strings.TrimPrefix(filename, "/")
}
