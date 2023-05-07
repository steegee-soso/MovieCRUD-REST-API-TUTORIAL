package Utils

import (
	"net/http"
	"regexp"
	"strings"
)

func validateEmail(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return pattern.MatchString(strings.TrimSpace(email))
}

func isEmptyField(field string) bool {
	return strings.TrimSpace(field) == ""
}

func isValidRequest(w http.ResponseWriter, r *http.Request, methodName string) bool {
	return r.Method == strings.ToUpper(methodName)
}
