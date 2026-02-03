package handlers

import (
   "net/mail"
   "regexp"
)

var (
    hasUpper   = regexp.MustCompile(`[A-Z]`)
    hasLower   = regexp.MustCompile(`[a-z]`)
    hasDigit   = regexp.MustCompile(`[0-9]`)
    hasSpecial = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func isValidEmail(email string) bool {
	_,err := mail.ParseAddress(email)
	return err == nil
}

func isValidPassword(p string) bool {
	if len(p)< 8 {
		return false
	}
	if !hasUpper.MatchString(p) {
        return false
    }
    if !hasLower.MatchString(p) {
        return false
    }
    if !hasDigit.MatchString(p) {
        return false
    }
    if !hasSpecial.MatchString(p) {
        return false
    }
    return true
}