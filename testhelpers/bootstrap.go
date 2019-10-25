package testhelpers

// BootstrapTestConfig will return a db instance
func BootstrapTestConfig() TestConfig {
	tc, err := InitTestConfig()
	if err != nil {
		panic(err.Error())
	}

	return tc
}
