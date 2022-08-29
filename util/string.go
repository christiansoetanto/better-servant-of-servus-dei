package util

import (
	"log"
	"regexp"
	"strings"
)

func Sanitize(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Println(err)
		return input
	}
	return strings.ToLower(reg.ReplaceAllString(input, ""))
}
