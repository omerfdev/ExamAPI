package utils

import "os"

func GetSwagHostEnv() string {
	e, ok := os.LookupEnv("SWAG_URL")
	if !ok {
		e = defaultSwagUrl
	}
	return e
}

func GetGoEnv() string {
	return defaultEnv
}


const defaultSwagUrl = "localhost:1993"
const defaultEnv = "local-qa"