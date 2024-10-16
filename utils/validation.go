package utils

import (
	"net/mail"

	"strconv"

	"regexp"
)
	

func ValidEmail(email string) bool {
  _, err := mail.ParseAddress(email)
  return err == nil
}

func IsStringNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func MatchDateTime(dateTime string) bool {
	matched, _ := regexp.MatchString(`\d{4}(.\d{2}){2}(\s|T)(\d{2}.){2}\d{2}`, dateTime)
	return matched;
}