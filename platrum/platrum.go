package platrum

import "net/url"

func CheckUserPassword(email, password string) bool {
	q := url.Values{}
	q.Set("email", email)
	q.Set("password", password)
	// @TODO add error logging
	isValid, _ := requestBool(, q)

	return isValid
}
