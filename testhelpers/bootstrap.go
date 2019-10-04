package testhelpers

// import (
// 	th "chitChat/testhelpers"
// )

// BootstrapTestConfig will return a db instance
func BootstrapTestConfig() TestConfig {
	tc, err := InitTestConfig()
	if err != nil {
		panic("Someting went wrong, check your environment variables")
	}

	return tc
}
