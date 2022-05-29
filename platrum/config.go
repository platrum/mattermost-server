package platrum

import "os"

const (
	envApiHost                  = "MESSENGER_API_HOST"
	envCallbackUrlPasswordCheck = "MESSENGER_CALLBACK_URL_PASSWORD_CHECK"
)

func apiHost() string {
	if res, ok := os.LookupEnv(envApiHost); ok {
		return res
	}

	return ""
}

func callbackUrlPasswordCheck() string {
	if res, ok := os.LookupEnv(envCallbackUrlPasswordCheck); ok {
		return res
	}

	return ""
}
