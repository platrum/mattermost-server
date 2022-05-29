package platrum

import "os"

const (
	envCoreHost = "MESSENGER_CORE_HOST"
	envCorePort = "MESSENGER_CORE_PORT"
)

func coreHost() string {
	if res, ok := os.LookupEnv(envCoreHost); ok {
		return res
	}

	return ""
}

func corePort() string {
	if res, ok := os.LookupEnv(envCorePort); ok {
		return res
	}

	return ""
}
